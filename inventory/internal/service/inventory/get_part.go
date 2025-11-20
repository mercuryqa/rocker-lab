package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (s *InventoryService) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	return s.repo.GetPart(ctx, uuid)
}
