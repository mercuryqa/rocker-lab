package repository

import (
	"context"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order model.OrderInfo) (string, error)
	PayOrder(ctx context.Context, id, status, paymentMethodName, transactionUuid string) bool
	GetOrder(ctx context.Context, id string) (*model.OrderInfo, bool)
	CancelOrder(ctx context.Context, id, status string) bool
	// CheckItemsIn() bool
}
