package converter

import (
	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PartsToProto(parts map[string]model.Part) map[string]*inventoryV1.Part {
	protoParts := map[string]*inventoryV1.Part{}
	for partUuid, part := range parts {
		protoParts[partUuid] = PartToProto(&part)
	}
	return protoParts
}

func PartToProto(part *model.Part) *inventoryV1.Part {
	if part == nil {
		return &inventoryV1.Part{}
	}
	return &inventoryV1.Part{
		Uuid:        part.UUID,
		Name:        part.Name,
		Description: part.Description, Price: part.Price,
		StockQuantity: part.StockQuantity,
		Category:      inventoryV1.Category(part.Category),
		Dimensions:    DimensionsToProto(&part.Dimensions),
		Manufacturer:  ManufacturerToProto(&part.Manufacturer),
		Tags:          part.Tags,
		CreatedAt:     timestamppb.New(part.CreatedAt),
		UpdatedAt:     timestamppb.New(part.UpdatedAt),
	}

}

func DimensionsToProto(dimensions *model.Dimensions) *inventoryV1.Dimensions {
	if dimensions == nil {
		return &inventoryV1.Dimensions{}
	}
	return &inventoryV1.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func ManufacturerToProto(manufacturer *model.Manufacturer) *inventoryV1.Manufacturer {
	if manufacturer == nil {
		return &inventoryV1.Manufacturer{}
	}
	return &inventoryV1.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func PartsFilterToModel(filter *inventoryV1.PartsFilter) model.PartsFilter {
	if filter == nil {
		return model.PartsFilter{}
	}
	return model.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            CategoriesToModel(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func CategoriesToModel(categories []inventoryV1.Category) []model.Category {
	if len(categories) == 0 {
		return nil
	}

	res := make([]model.Category, len(categories))
	for i, c := range categories {
		res[i] = model.Category(c)
	}
	return res
}

func GetPartRequestToProto(req *model.GetPartRequest) *inventoryV1.GetPartRequest {
	return &inventoryV1.GetPartRequest{
		InventoryUuid: req.InventoryUuid,
	}
}

func GetPartResponseToModel(resp *inventoryV1.GetPartResponse) *model.GetPartResponse {
	if resp == nil {
		return nil
	}
	return &model.GetPartResponse{
		Part: model.Part{
			UUID:          resp.Part.Uuid,
			Name:          resp.Part.Name,
			Description:   resp.Part.Description,
			Price:         resp.Part.Price,
			StockQuantity: resp.Part.StockQuantity,
			Category:      model.Category(resp.Part.Category), // если есть enum-конвертер, используем его
			// остальные поля
		},
	}
}

func PartsFilterToProto(filter model.PartsFilter) *inventoryV1.GetListPartRequest {
	return &inventoryV1.GetListPartRequest{
		Filter: &inventoryV1.PartsFilter{
			Uuids: filter.Uuids,
			Names: filter.Names,
			Tags:  filter.Tags,
			// Categories нужно конвертировать в protobuf enum
			Categories: CategoriesToProto(filter.Categories),
		},
	}
}

func CategoriesToProto(cats []model.Category) []inventoryV1.Category {
	res := make([]inventoryV1.Category, len(cats))
	for i, c := range cats {
		res[i] = CategoryToProto(c)
	}
	return res
}

func PartFromProto(p *inventoryV1.Part) model.Part {
	return model.Part{
		UUID:          p.Uuid,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      model.Category(p.Category), // при необходимости через CategoryFromProto
		Tags:          p.Tags,
		// и Dimensions, Manufacturer, Metadata
	}
}

// CategoryToProto конвертирует model.Category в inventory_v1.Category
func CategoryToProto(c model.Category) inventoryV1.Category {
	switch c {
	case model.CategoryEngine:
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
