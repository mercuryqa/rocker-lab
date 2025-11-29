package service

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *model.GetOrderResponse) error
	PayOrder(ctx context.Context, id, status, paymentMethodName, transactionUuid string) bool
	GetOrder(ctx context.Context, id string) (*model.GetOrderResponse, bool)
	CancelOrder(ctx context.Context, id, status string) bool
	CheckItems(ctx context.Context, ids []string) bool
	// CheckItemsIn() bool
}
