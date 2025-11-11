package inventory

import (
	"context"

	converter "github.com/mercuryqa/rocket-lab/inventory/internal/converter"
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
	protoReq := converter.GetPartRequestToProto(info) // model → protobuf
	protoResp, err := s.repo.GetPart(ctx, protoReq)   // вызов репозитория
	if err != nil {
		return nil, err
	}
	return converter.GetPartResponseToModel(protoResp), nil // protobuf → model
}

func (s *InventoryService) ListParts(ctx context.Context, info model.PartsFilter) (*model.ListPartsResponse, error) {
	protoReq := converter.PartsFilterToProto(info)    // model → protobuf
	protoResp, err := s.repo.ListParts(ctx, protoReq) // вызываем репозиторий
	if err != nil {
		return nil, err
	}

	parts := make([]model.Part, len(protoResp.Parts))
	for i, p := range protoResp.Parts {
		parts[i] = converter.PartFromProto(p) // конвертер protobuf → model.Part
	}

	return &model.ListPartsResponse{Parts: parts}, nil
}
