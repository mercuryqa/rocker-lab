package inventory

import (
	"context"

	domain "github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository/converter"
)

func (r *InventoryRepository) ListParts(ctx context.Context, filter domain.PartsFilter) (domain.ListPartsResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Part, 0, len(r.inventory)) // возвращаем domain.Part

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
				if string(repoPart.Category) == string(c) { // сравниваем по строковому значению
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// конвертируем из repository.model.Part в domain.model.Part
		domainPart := converter.RepoPartToDomain(repoPart)
		result = append(result, domainPart)
	}

	return domain.ListPartsResponse{Parts: result}, nil
}
