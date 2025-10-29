package order

func (s *service) PayOrder(id, status string) bool {
	s.orderRepository.PayOrder(id, status)
	return true
}
