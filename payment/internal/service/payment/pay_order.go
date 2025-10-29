package payment

import (
	"context"

	"github.com/mercuryqa/rocket-lab/payment/internal/model"
)

func (s *service) PayOrder(ctx context.Context, info model.PayOrderRequest) (model.PayOrderResponse, error) {
	// дополнительная логика
	transactionUuid, err := s.paymentRepository.PayOrder(ctx, info)
	if err != nil {
		return model.PayOrderResponse{}, err
	}
	return transactionUuid, nil
}
