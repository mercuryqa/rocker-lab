package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository/converter"
)

func (r *InventoryRepository) GetPart(_ context.Context, uuid string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.inventory[uuid]
	if !ok {
		return model.Part{}, ErrNotFound
	}

	return converter.RepoModelToModel(part), nil
}
