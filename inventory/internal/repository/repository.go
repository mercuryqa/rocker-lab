package repository

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, info model.GetPartRequest) (model.GetPartResponse, error)
	ListParts(ctx context.Context, info model.ListPartsRequest) (model.ListPartsResponse, error)
}
