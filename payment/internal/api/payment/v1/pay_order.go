package api_v1

import (
	"context"

	"github.com/mercuryqa/rocket-lab/payment/internal/converter"
	"github.com/mercuryqa/rocket-lab/payment/internal/model"
	"github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
)

func (a *api) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	// Конвертируем proto → internal model
	info := model.PayOrderRequest{
		OrderUuid:     req.GetOrderUuid(),
		UserUuid:      req.GetUserUuid(),
		PaymentMethod: converter.ToModelPaymentMethod(req.GetPaymentMethod()),
	}

	resp, err := a.paymentService.PayOrder(ctx, info)
	if err != nil {
		return nil, err
	}

	// Конвертируем internal model → proto
	return &payment_v1.PayOrderResponse{
		TransactionUuid: resp.TransactionUuid,
	}, nil
}
