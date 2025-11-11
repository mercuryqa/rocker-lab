package inventory

import (
	"context"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

// Публичный метод реализует gRPC интерфейс
func (s *InventoryRepository) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	partModel, ok := s.inventory[req.InventoryUuid] // тут у тебя map[string]*model.GetPartResponse
	if !ok {
		return nil, ErrNotFound
	}

	// конвертация model → proto
	return &inventoryV1.GetPartResponse{
		Part: partModel.Part,
	}, nil
}
