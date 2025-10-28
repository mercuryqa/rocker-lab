package service

import (
	"context"

	"github.com/mercuryqa/rocket-lab/payment/internal/model"
)

type PaymentService interface {
	PayOrder(ctx context.Context, info model.PayOrderRequest) (model.PayOrderResponse, error)
}
