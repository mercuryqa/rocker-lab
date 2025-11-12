package converter

import (
	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	repoModel "github.com/mercuryqa/rocket-lab/inventory/internal/repository/model"
)

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

//func PartsToModel(parts map[string]repoModel.GetPartResponse) map[string]model.Part {
//	modelParts := map[string]model.Part{}
//	for partUuid, part := range parts {
//		modelParts[partUuid] = repoModelToModel(part)
//	}
//	return modelParts
//}

func RepoModelToModel(part *repoModel.GetPartResponse) model.Part {
	return model.Part{
		UUID:          part.Part.UUID,
		Name:          part.Part.Name,
		Description:   part.Part.Description,
		Price:         part.Part.Price,
		StockQuantity: part.Part.StockQuantity,
		Category:      model.Category(part.Part.Category),
		Dimensions:    DimensionsToModel(&part.Part.Dimensions),
		Manufacturer:  ManufacturerToModel(&part.Part.Manufacturer),
		Tags:          part.Part.Tags,
		CreatedAt:     part.Part.CreatedAt,
		UpdatedAt:     part.Part.UpdatedAt,
	}
}
