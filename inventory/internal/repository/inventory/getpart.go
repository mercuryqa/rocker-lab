package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (r *InventoryRepository) GetPart(ctx context.Context, req *model.GetPartRequest) (*model.GetPartResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.inventory[req.InventoryUuid]
	if !ok {
		return nil, ErrNotFound
	}
	return part, nil
}
