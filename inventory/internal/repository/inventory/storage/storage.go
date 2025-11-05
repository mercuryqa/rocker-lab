package storage

import (
	"context"
	"errors"
	"strings"
	"sync"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

var ErrNotFound = errors.New("part not found")

// InventoryStorage представляет потокобезопасное хранилище данных о заказах
type InventoryStorage struct {
	inventoryV1.UnimplementedInventoryStorageServer

	mu        sync.Mutex
	inventory map[string]*inventoryV1.GetPartResponse
}

func NewInventoryStorage() *InventoryStorage {
	s := &InventoryStorage{
		inventory: make(map[string]*inventoryV1.GetPartResponse),
	}
	GenerateSampleData(s)
	return s
}

// Публичный метод реализует gRPC интерфейс
func (s *InventoryStorage) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	part, ok := s.inventory[req.GetInventoryUuid()]
	if !ok {
		return nil, ErrNotFound
	}
	return part, nil
}

// helper to check if filter is empty
func isFilterEmpty(f *inventoryV1.PartsFilter) bool {
	if f == nil {
		return true
	}
	return len(f.Uuids) == 0 && len(f.Names) == 0 && len(f.Categories) == 0 && len(f.ManufacturerCountries) == 0 && len(f.Tags) == 0
}

// ListParts implements filtering logic described in the task
func (s *InventoryStorage) ListParts(ctx context.Context, req *inventoryV1.GetListPartRequest) (*inventoryV1.GetListPartResponse, error) {
	parts := make([]*inventoryV1.Part, 0, len(s.inventory))
	for _, pResp := range s.inventory {
		parts = append(parts, pResp.Part) // взять Part из GetPartResponse
	}
	f := req.GetFilter()
	if isFilterEmpty(f) {
		return &inventoryV1.GetListPartResponse{Parts: parts}, nil
	}

	var err error
	if len(f.Uuids) > 0 {
		parts, err = filterByUUIDs(parts, f.Uuids)
		if err != nil {
			return nil, err
		}
	}
	if len(f.Names) > 0 {
		parts = filterByNames(parts, f.Names)
	}
	if len(f.Categories) > 0 {
		// Преобразуем []string в []inventoryV1.Category
		categoriesEnums := make([]inventoryV1.Category, 0, len(f.Categories))
		for _, c := range f.Categories {
			cat := inventoryV1.Category(inventoryV1.Category_value[strings.ToUpper(c)])
			categoriesEnums = append(categoriesEnums, cat)
		}
		parts = filterByCategories(parts, categoriesEnums)
	}
	if len(f.ManufacturerCountries) > 0 {
		parts = filterByCountries(parts, f.ManufacturerCountries)
	}
	if len(f.Tags) > 0 {
		parts = filterByTags(parts, f.Tags)
	}
	return &inventoryV1.GetListPartResponse{Parts: parts}, nil
}

// Filter helpers
func filterByUUIDs(parts []*inventoryV1.Part, uuids []string) ([]*inventoryV1.Part, error) {
	set := make(map[string]struct{}, len(uuids))
	for _, u := range uuids {
		if strings.TrimSpace(u) == "" {
			continue
		}
		set[u] = struct{}{}
	}
	if len(set) == 0 {
		return nil, errors.New("provided uuids are empty")
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if _, ok := set[p.Uuid]; ok {
			out = append(out, p)
		}
	}
	return out, nil
}

func filterByNames(parts []*inventoryV1.Part, names []string) []*inventoryV1.Part {
	set := make(map[string]struct{}, len(names))
	for _, n := range names {
		set[n] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if _, ok := set[p.Name]; ok {
			out = append(out, p)
		}
	}
	return out
}

func filterByCategories(parts []*inventoryV1.Part, cats []inventoryV1.Category) []*inventoryV1.Part {
	set := make(map[inventoryV1.Category]struct{}, len(cats))
	for _, c := range cats {
		set[c] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if _, ok := set[p.Category]; ok {
			out = append(out, p)
		}
	}
	return out
}

func filterByCountries(parts []*inventoryV1.Part, countries []string) []*inventoryV1.Part {
	set := make(map[string]struct{}, len(countries))
	for _, c := range countries {
		set[c] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if p.Manufacturer != nil {
			if _, ok := set[p.Manufacturer.Country]; ok {
				out = append(out, p)
			}
		}
	}
	return out
}

func filterByTags(parts []*inventoryV1.Part, tags []string) []*inventoryV1.Part {
	set := make(map[string]struct{}, len(tags))
	for _, t := range tags {
		set[t] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		for _, pt := range p.Tags {
			if _, ok := set[pt]; ok {
				out = append(out, p)
				break
			}
		}
	}
	return out
}
