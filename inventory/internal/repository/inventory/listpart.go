package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (r *InventoryRepository) ListParts(ctx context.Context, filter model.PartsFilter) (model.ListPartsResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Part, 0, len(r.inventory)) // теперь []Part, не []*Part

	for _, getPartResp := range r.inventory {
		part := getPartResp.Part // model.Part

		// фильтр по UUID
		if len(filter.Uuids) > 0 {
			found := false
			for _, u := range filter.Uuids {
				if part.UUID == u {
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
				if part.Name == n {
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
				if part.Category == c {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		result = append(result, part) // append копию структуры, не указатель
	}

	return model.ListPartsResponse{Parts: result}, nil
}
