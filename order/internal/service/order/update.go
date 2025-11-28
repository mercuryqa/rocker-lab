package order

import "github.com/mercuryqa/rocket-lab/order/internal/model"

func (s *service) UpdateOrder(id string, paymentMethod model.PaymentMethod, transactionUuid string) bool {
	s.orderRepository.UpdateOrder(id, paymentMethod, transactionUuid)
	return true
}
