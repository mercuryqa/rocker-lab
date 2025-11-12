package apiv1

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/converter"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.GetPart(ctx, req.GetInventoryUuid())
	if err != nil {
		return nil, err
	}
	return &inventoryV1.GetPartResponse{
		Part: converter.PartToProto(&part.Part),
	}, nil
}
