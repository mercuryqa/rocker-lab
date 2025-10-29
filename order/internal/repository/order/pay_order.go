package order

func (s *OrderRepository) PayOrder(id, status string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[id]
	if !ok {
		return false
	}

	order.Status = status
	return true
}
