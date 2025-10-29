package repository

import (
	"context"

	"github.com/mercuryqa/rocket-lab/payment/internal/model"
)

type PaymentRepository interface {
	PayOrder(ctx context.Context, info model.PayOrderRequest) (model.PayOrderResponse, error)
}
