package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func ToProtoParts(parts []model.Part) []*inventoryV1.Part {
	protoList := make([]*inventoryV1.Part, len(parts))

	for i := range parts {
		protoList[i] = ToProtoPart(&parts[i])
	}

	return protoList
}

func ToProtoPart(part *model.Part) *inventoryV1.Part {
	if part == nil {
		return &inventoryV1.Part{}
	}
	return &inventoryV1.Part{
		Uuid:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      ToProtoCategory(part.Category),
		Dimensions:    ToProtoDimensions(&part.Dimensions),
		Manufacturer:  ToProtoManufacturer(&part.Manufacturer),
		Tags:          part.Tags,
		CreatedAt:     timestamppb.New(part.CreatedAt),
		UpdatedAt:     timestamppb.New(part.UpdatedAt),
	}
}

func ToProtoDimensions(dimensions *model.Dimensions) *inventoryV1.Dimensions {
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

func ToProtoManufacturer(manufacturer *model.Manufacturer) *inventoryV1.Manufacturer {
	if manufacturer == nil {
		return &inventoryV1.Manufacturer{}
	}
	return &inventoryV1.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func ToModelPartsFilter(filter *inventoryV1.PartsFilter) model.PartsFilter {
	if filter == nil {
		return model.PartsFilter{}
	}
	return model.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            ToModelCategories(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func ToModelCategories(categories []inventoryV1.Category) []model.Category {
	if len(categories) == 0 {
		return nil
	}

	res := make([]model.Category, len(categories))
	for i, c := range categories {
		res[i] = model.Category(c)
	}
	return res
}

func ToProtoGetPartRequest(req *model.GetPartRequest) *inventoryV1.GetPartRequest {
	return &inventoryV1.GetPartRequest{
		InventoryUuid: req.InventoryUuid,
	}
}

func ToProtoPartsFilter(filter model.PartsFilter) *inventoryV1.GetListPartRequest {
	return &inventoryV1.GetListPartRequest{
		Filter: &inventoryV1.PartsFilter{
			Uuids: filter.Uuids,
			Names: filter.Names,
			Tags:  filter.Tags,
			// Categories нужно конвертировать в protobuf enum
			Categories: ToProtoCategories(filter.Categories),
		},
	}
}

func ToProtoCategories(cats []model.Category) []inventoryV1.Category {
	res := make([]inventoryV1.Category, len(cats))
	for i, c := range cats {
		res[i] = ToProtoCategory(c)
	}
	return res
}

func ToPartProto(p *inventoryV1.Part) model.Part {
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

// ToProtoCategory конвертирует model.Category в inventory_v1.Category
func ToProtoCategory(c model.Category) inventoryV1.Category {
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
