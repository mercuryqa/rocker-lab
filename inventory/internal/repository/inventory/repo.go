package inventory

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

var ErrNotFound = errors.New("part not found")

var _ def.InventoryRepository = (*InventoryRepository)(nil)

// InventoryStorage представляет потокобезопасное хранилище данных о заказах
type InventoryRepository struct {
	inventoryV1.UnimplementedInventoryStorageServer

	mu        sync.RWMutex
	inventory map[string]*inventoryV1.GetPartResponse
}

func NewInventoryRepository() *InventoryRepository {
	s := &InventoryRepository{
		inventory: make(map[string]*inventoryV1.GetPartResponse),
	}
	GenerateSampleData(s)
	return s
}

// Публичный метод реализует gRPC интерфейс
func (s *InventoryRepository) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	partModel, ok := s.inventory[req.InventoryUuid] // тут у тебя map[string]*model.GetPartResponse
	if !ok {
		return nil, ErrNotFound
	}

	// конвертация model → proto
	return &inventoryV1.GetPartResponse{
		Part: partModel.Part,
	}, nil
}

// helper to check if filter is empty
func isFilterEmpty(f *inventoryV1.PartsFilter) bool {
	if f == nil {
		return true
	}
	return len(f.Uuids) == 0 && len(f.Names) == 0 && len(f.Categories) == 0 && len(f.ManufacturerCountries) == 0 && len(f.Tags) == 0
}

// CategoryToProto конвертирует модельную категорию в protobuf
func CategoryToProto(c model.Category) inventoryV1.Category {
	switch c {
	case model.CategoryUnknown:
		return inventoryV1.Category_ENGINE
	case model.CategoryFuel:
		return inventoryV1.Category_FUEL
	case model.CategoryPorthole:
		return inventoryV1.Category_PORTHOLE
	case model.CategoryWing:
		return inventoryV1.Category_WING
	default:
		return inventoryV1.Category_UNKNOWN
	}
}

// ListParts возвращает список деталей с фильтрацией
func (s *InventoryRepository) ListParts(ctx context.Context, req *inventoryV1.GetListPartRequest) (*inventoryV1.GetListPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Сначала создаём protobuf Part из модели
	parts := make([]*inventoryV1.Part, 0, len(s.inventory))
	for _, pResp := range s.inventory {
		part := &inventoryV1.Part{
			Uuid:          pResp.Part.Uuid,
			Name:          pResp.Part.Name,
			Description:   pResp.Part.Description,
			Price:         pResp.Part.Price,
			StockQuantity: pResp.Part.StockQuantity,
			Category:      pResp.Part.Category,
			Dimensions:    pResp.Part.Dimensions,
			Manufacturer:  pResp.Part.Manufacturer,
			Tags:          pResp.Part.Tags,
		}
		parts = append(parts, part)
	}

	// Фильтр
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
		categoriesEnums := make([]inventoryV1.Category, 0, len(f.Categories))
		for _, c := range f.Categories {
			categoriesEnums = append(categoriesEnums, CategoryToProto(model.Category(c)))
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
