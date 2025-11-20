package order

import "github.com/mercuryqa/rocket-lab/order/internal/model"

func (s *service) CancelOrder(id string, status model.OrderStatus) bool {
	s.orderRepository.CancelOrder(id, status)
	return true
}
