package order

import "github.com/mercuryqa/rocket-lab/order/model"

func (s *service) GetOrder(id string) (*model.GetOrderResponse, bool) {
	order, ok := s.orderRepository.GetOrder(id)
	if !ok {
		return nil, false
	}
	return order, true
}
