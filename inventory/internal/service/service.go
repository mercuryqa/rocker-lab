package service

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, req *model.GetPartRequest) (*model.GetPartResponse, error)
	ListParts(ctx context.Context, filter model.PartsFilter) (*model.ListPartsResponse, error)
}
