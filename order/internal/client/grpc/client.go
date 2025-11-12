package grpc

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUuid, userUuid string, paymetnMethod paymentV1.PaymentMethod) (string, error)
}
