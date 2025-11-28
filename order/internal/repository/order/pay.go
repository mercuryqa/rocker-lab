package order

import (
	"github.com/mercuryqa/rocket-lab/order/internal/converter"
	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (r *OrderRepository) PayOrder(id string, status model.OrderStatus) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[id]
	if !ok {
		return false
	}

	order.Status = converter.ToRepoModelModel(status)
	return true
}
