package service

import (
	"github.com/mercuryqa/rocket-lab/order/model"
)

type OrderService interface {
	CreateOrder(order *model.GetOrderResponse)
	PayOrder(id, paymentMethod string) string
	GetOrder(id string) (*model.GetOrderResponse, bool)
	CancelOrder(id, status string) bool
	UpdateOrder(id, paymentMethod, transactionUuid string) bool
}
