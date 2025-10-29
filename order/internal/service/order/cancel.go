package order

func (s *service) CancelOrder(id, status string) bool {
	s.orderRepository.CancelOrder(id, status)
	return true
}
