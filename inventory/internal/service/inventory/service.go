package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/service"
)

var _ def.InventoryService = (*InventoryService)(nil)

type InventoryService struct {
	repo repository.InventoryRepository // repo работает с model
}

func NewService(repo repository.InventoryRepository) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

func (s *InventoryService) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	return s.repo.GetPart(ctx, uuid)
}

func (s *InventoryService) ListParts(ctx context.Context, filter model.PartsFilter) (*model.ListPartsResponse, error) {
	resp, err := s.repo.ListParts(ctx, filter) // resp — model.ListPartsResponse
	if err != nil {
		return nil, err
	}
	return &resp, nil // берем адрес структуры, чтобы получить *model.ListPartsResponse
}
