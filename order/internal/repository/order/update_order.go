package order

func (s *OrderService) UpdateOrder(id, paymentMethod, transactionUuid string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[id]
	if !ok {
		return false
	}

	order.TransactionUuid = transactionUuid
	order.PaymentMethod = paymentMethod

	return true
}
