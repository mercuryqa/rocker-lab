package service

import (
	"github.com/mercuryqa/rocket-lab/order/model"
)

type OrderService interface {
	CreateOrder(order *model.GetOrderResponse)
	PayOrder(id, status string) bool
	GetOrder(id string) (*model.GetOrderResponse, bool)
	CanselOrder(id, status string) bool
	UpdateOrder(id, paymentMethod, transactionUuid string) bool
}
