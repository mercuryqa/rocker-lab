package inventory

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	"github.com/mercuryqa/rocket-lab/inventory/internal/repository"
	def "github.com/mercuryqa/rocket-lab/inventory/internal/service"
)

var _ def.InventoryService = (*InventoryService)(nil)

type InventoryService struct {
	inventoryRepository repository.InventoryRepository
}

func NewService(inventoryRepository repository.InventoryRepository) *InventoryService {
	return &InventoryService{
		inventoryRepository: inventoryRepository,
	}
}

func (s *InventoryService) GetPart(ctx context.Context, info *model.GetPartRequest) (*model.GetPartResponse, error) {
	part, err := s.inventoryRepository.GetPart(ctx, info)
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
