package order

import "github.com/mercuryqa/rocket-lab/order/model"

// CreateOrder создает заказ
func (s *OrderRepository) CreateOrder(order *model.GetOrderResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders[order.OrderUuid] = order
}
