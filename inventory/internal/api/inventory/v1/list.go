package apiv1

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/converter"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.GetListPartRequest) (*inventoryV1.GetListPartResponse, error) {
	// Внутри берем фильтр
	filter := converter.PartsFilterToModel(req.GetFilter())

	list, err := a.inventoryService.ListParts(ctx, filter)
	if err != nil {
		return &inventoryV1.GetListPartResponse{}, err
	}

	protoParts := make([]*inventoryV1.Part, len(list.Parts))
	for i, p := range list.Parts {
		protoParts[i] = converter.PartToProto(&p)
	}

	return &inventoryV1.GetListPartResponse{
		Parts: protoParts,
	}, nil
}
