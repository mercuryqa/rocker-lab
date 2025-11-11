package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/service"
)

var _ def.InventoryService = (*InventoryService)(nil)

type InventoryService struct {
	repo repository.InventoryRepository
}

func NewService(repo repository.InventoryRepository) *InventoryService {
	return &InventoryService{
		repo: repo,
	}
}

func (s *InventoryService) GetPart(ctx context.Context, info *model.GetPartRequest) (*model.GetPartResponse, error) {
	part, err := s.repo.GetPart(ctx, info)
	if err != nil {
		return &model.GetPartResponse{}, err
	}
	return part, nil
}

// func (s *InventoryService) ListPart(ctx context.Context, info model.ListPartsRequest) (model.ListPartsResponse, error) {
//	list, err := s.inventoryRepository.ListParts(ctx, info)
//	if err != nil {
//		return model.ListPartsResponse{}, err
//	}
//	return list, nil
//}
