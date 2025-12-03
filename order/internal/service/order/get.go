package order

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (s *service) GetOrder(ctx context.Context, id string) (*model.OrderInfo, bool) {
	order, ok := s.orderRepository.GetOrder(ctx, id)
	if !ok {
		return nil, false
	}
	return order, true
}
