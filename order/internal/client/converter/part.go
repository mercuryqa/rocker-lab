package converter

import (
	"github.com/mercuryqa/rocket-lab/order/internal/model"

	generatedInventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func PartListToModel(parts []*generatedInventoryV1.Part) []model.Part {
	result := make([]model.Part, 0, len(parts))
	for _, p := range parts {
		result = append(result, PartToModel(p))
	}
	return result
}

// Конвертация одной детали
func PartToModel(p *generatedInventoryV1.Part) model.Part {
	return model.Part{
		UUID:          p.Uuid,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      model.Category(p.Category),
		Dimensions: model.Dimensions{
			Width:  p.Dimensions.Width,
			Height: p.Dimensions.Height,
			Length: p.Dimensions.Length,
			Weight: p.Dimensions.Weight,
		},
		Manufacturer: model.Manufacturer{
			Name:    p.Manufacturer.Name,
			Country: p.Manufacturer.Country,
		},
		Tags: p.Tags,
	}
}

func PartsFilterToProto(f model.PartsFilter) *generatedInventoryV1.PartsFilter {
	return &generatedInventoryV1.PartsFilter{
		Uuids:                 f.Uuids,
		Names:                 f.Names,
		Categories:            CategoriesToProto(f.Categories),
		ManufacturerCountries: f.ManufacturerCountries,
		Tags:                  f.Tags,
	}
}

// вспомогательная функция для категорий
func CategoriesToProto(categories []model.Category) []generatedInventoryV1.Category {
	result := make([]generatedInventoryV1.Category, 0, len(categories))
	for _, c := range categories {
		result = append(result, generatedInventoryV1.Category(c))
	}
	return result
}
