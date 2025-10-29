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

	// Генерируем новый UUID для транзакции
	transactionUUID := uuid.NewString()

	log.Printf("Создана транзакция %s для заказа %s", transactionUUID, info.OrderUuid)

	// Возвращаем ответ клиенту
	return model.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil

	// resp := model.PayOrderResponse{
	//	TransactionUuid: transactionUUID,
	// }
	// r.data[info.OrderUuid] = resp
	// return resp, nil
}
