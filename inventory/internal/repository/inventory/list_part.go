package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository/converter"
	repoModel "github.com/mercuryqa/rocket-lab/inventory/internal/repository/model"
)

func (r *InventoryRepository) ListParts(ctx context.Context, filter model.PartsFilter) (model.ListPartsResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	partsFiltered := make([]model.Part, 0, len(r.inventory)) // возвращаем domain.Part

	for _, getPartResp := range r.inventory {
		repoPart := getPartResp.Part // repoModel.Part

		// фильтр по UUID
		if len(filter.Uuids) > 0 {
			found := false
			for _, u := range filter.Uuids {
				if repoPart.UUID == u {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// фильтр по имени
		if len(filter.Names) > 0 {
			found := false
			for _, n := range filter.Names {
				if repoPart.Name == n {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// фильтр по категориям
		if len(filter.Categories) > 0 {
			found := false
			for _, c := range filter.Categories {
				// сравнение через явное приведение к одному типу
				if repoPart.Category == repoModel.Category(c) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// конвертируем из repository.model.Part в domain.model.Part
		domainPart := converter.RepoPartToModel(repoPart)
		partsFiltered = append(partsFiltered, domainPart)
	}

	return model.ListPartsResponse{Parts: partsFiltered}, nil
}
