package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (s *InventoryService) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	parts, errRep := s.repo.ListParts(ctx, filter)
	if errRep != nil {
		return []model.Part{}, errRep
	}
	return parts, nil
}
