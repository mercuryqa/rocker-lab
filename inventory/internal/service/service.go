package service

import (
	"context"

	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, string2 string) (model.Part, error)
	ListParts(ctx context.Context, filter model.PartsFilter) (*model.ListPartsResponse, error)
}
