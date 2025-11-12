package grpc

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUuid, userUuid, paymetnMethod string)
}
