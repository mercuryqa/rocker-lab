package order

import (
	"context"
)

func (s *service) PayOrder(ctx context.Context, id, status, paymentMethodName, transactionUuid string) bool {
	s.orderRepository.PayOrder(ctx, id, status, paymentMethodName, transactionUuid)
	return true
}
