package apiv1

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/converter"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.GetListPartRequest) (*inventoryV1.GetListPartResponse, error) {
	reqFilter := req.GetFilter()

	partsResp, err := a.inventoryService.ListParts(ctx, converter.ToModelPartsFilter(reqFilter))
	if err != nil {
		return &inventoryV1.GetListPartResponse{}, err
	}

	return &inventoryV1.GetListPartResponse{
		Parts: converter.ToProtoParts(partsResp), // <- исправлено здесь
	}, nil
}
