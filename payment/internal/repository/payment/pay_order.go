package payment

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/mercuryqa/rocket-lab/payment/internal/model"
)

func (r *repository) PayOrder(_ context.Context, info model.PayOrderRequest) (model.PayOrderResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	transactionUUID := uuid.NewString()

	log.Printf("Создана транзакция %s для заказа %s", transactionUUID, info.OrderUuid)

	return model.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}
