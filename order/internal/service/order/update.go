package order

func (s *service) UpdateOrder(id, paymentMethod, transactionUuid string) bool {
	s.orderRepository.UpdateOrder(id, paymentMethod, transactionUuid)
	return true
}
