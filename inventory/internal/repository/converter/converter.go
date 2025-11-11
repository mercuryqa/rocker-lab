package converter

import (
	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	repoModel "github.com/mercuryqa/rocket-lab/inventory/internal/repository/model"
)

func PartsToModel(parts map[string]repoModel.Part) map[string]model.Part {
	modelParts := map[string]model.Part{}
	for partUuid, part := range parts {
		modelParts[partUuid] = PartToModel(part)
	}
	return modelParts
}

func PartToModel(part repoModel.Part) model.Part {
	return model.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model.Category(part.Category),
		Dimensions:    DimensionsToModel(part.Dimensions),
		Manufacturer:  ManufacturerToModel(part.Manufacturer),
		Tags:          part.Tags,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}
}

func DimensionsToModel(dimensions *repoModel.Dimensions) model.Dimensions {
	if dimensions == nil {
		return model.Dimensions{}
	}
	return model.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func ManufacturerToModel(manufacturer *repoModel.Manufacturer) model.Manufacturer {
	if manufacturer == nil {
		return model.Manufacturer{}
	}
	return model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}
