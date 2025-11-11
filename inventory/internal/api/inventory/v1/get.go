package apiv1

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/converter"
	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

// api/apiv1/grpc.go
func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	modelReq := &model.GetPartRequest{InventoryUuid: req.InventoryUuid}
	part, err := a.inventoryService.GetPart(ctx, modelReq)
	if err != nil {
		return nil, err
	}
	return &inventoryV1.GetPartResponse{
		Part: converter.PartToProto(&part.Part),
	}, nil
}
