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
