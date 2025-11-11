package repository

import (
	"context"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error)
	//ListParts(ctx context.Context, info *inventoryV1.GetListPartRequest) (*inventoryV1.GetListPartResponse, error)
}
