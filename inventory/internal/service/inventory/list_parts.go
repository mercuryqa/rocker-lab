package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

func (s *InventoryService) ListParts(ctx context.Context, filter model.PartsFilter) (*model.ListPartsResponse, error) {
	resp, err := s.repo.ListParts(ctx, filter) // resp — model.ListPartsResponse
	if err != nil {
		return nil, err
	}
	return &resp, nil // берем адрес структуры, чтобы получить *model.ListPartsResponse
}
