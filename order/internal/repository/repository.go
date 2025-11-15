package repository

import (
	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) error
	PayOrder(id string, status model.OrderStatus) bool
	GetOrder(id string) (*model.Order, bool)
	CancelOrder(id string, status model.OrderStatus) bool
	UpdateOrder(id string, paymentMethod model.PaymentMethod, transactionUuid string) bool
}
