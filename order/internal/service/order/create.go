package order

import "github.com/mercuryqa/rocket-lab/order/model"

// CreateOrder создает заказ
func (s *service) CreateOrder(order *model.GetOrderResponse) {
	s.orderRepository.CreateOrder(order)
}
