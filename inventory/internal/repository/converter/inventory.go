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

func RepoPartToModel(p repoModel.Part) model.Part {
	return model.Part{
		UUID:          p.UUID,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      model.Category(p.Category),
		Dimensions:    model.Dimensions(p.Dimensions),
		Manufacturer:  model.Manufacturer(p.Manufacturer),
		Tags:          p.Tags,
		// Metadata:      p.Metadata,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ModelPartToRepo(p model.Part) repoModel.Part {
	return repoModel.Part{
		UUID:          p.UUID,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      repoModel.Category(p.Category),
		Dimensions:    repoModel.Dimensions(p.Dimensions),
		Manufacturer:  repoModel.Manufacturer(p.Manufacturer),
		Tags:          p.Tags,
		// Metadata:      p.Metadata,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
