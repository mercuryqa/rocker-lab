package service

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, info *model.GetPartRequest) (*model.GetPartResponse, error)
	ListParts(ctx context.Context, info model.PartsFilter) (*model.ListPartsResponse, error)
}
