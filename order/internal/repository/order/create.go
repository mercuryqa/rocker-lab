package order

import (
	"github.com/mercuryqa/rocket-lab/order/internal/model"
	repoModel "github.com/mercuryqa/rocket-lab/order/internal/repository/model"
)

// CreateOrder создает заказ
func (r *OrderRepository) CreateOrder(order *model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderSave := &repoModel.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: "",
		PaymentMethod:   "",
		Status:          repoModel.PendingPayment,
	}

	r.orders[order.OrderUuid] = orderSave

	return nil
}
