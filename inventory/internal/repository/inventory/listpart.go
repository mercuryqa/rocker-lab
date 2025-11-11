package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

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
