package order

import "github.com/mercuryqa/rocket-lab/order/internal/model"

func (s *service) GetOrder(id string) (*model.Order, bool) {
	order, ok := s.orderRepository.GetOrder(id)
	if !ok {
		return nil, false
	}
	return order, true
}
