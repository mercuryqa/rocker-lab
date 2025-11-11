package repository

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, info *model.GetPartRequest) (*model.GetPartResponse, error)
	ListParts(ctx context.Context, info *inventoryV1.GetListPartRequest) (*inventoryV1.GetListPartResponse, error)
}
