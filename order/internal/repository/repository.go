package repository

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.GetOrderResponse) (string, error)
	PayOrder(ctx context.Context, id, status, paymentMethodName, transactionUuid string) bool
	GetOrder(ctx context.Context, id string) (*model.GetOrderResponse, bool)
	CancelOrder(ctx context.Context, id, status string) bool
	// CheckItemsIn() bool
	CheckItems(ctx context.Context, ids []string) bool
}
