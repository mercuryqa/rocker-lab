package order

import (
	"github.com/mercuryqa/rocket-lab/order/model"
)

func (s *OrderService) GetOrder(id string) (*model.GetOrderResponse, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[id]
	if !ok {
		return nil, false
	}

	return order, true
}
