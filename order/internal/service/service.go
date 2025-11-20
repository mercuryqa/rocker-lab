package service

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *model.OrderRequest) (*model.OrderResponse, error)
	PayOrder(id string, paymentMethod model.PaymentMethod) string
	GetOrder(id string) (*model.Order, bool)
	CancelOrder(id string, status model.OrderStatus) bool
	UpdateOrder(id string, paymentMethod model.PaymentMethod, transactionUuid string) bool
}
