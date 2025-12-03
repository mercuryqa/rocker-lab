package order

import (
	"context"
)

func (s *service) CancelOrder(ctx context.Context, id, status string) bool {
	s.orderRepository.CancelOrder(ctx, id, status)
	return true
}
