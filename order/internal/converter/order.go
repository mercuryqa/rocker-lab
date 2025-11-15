package converter

import (
	"github.com/mercuryqa/rocket-lab/order/internal/model"
	repoModel "github.com/mercuryqa/rocket-lab/order/internal/repository/model"
)

func RepoModelToModel(order *repoModel.Order) *model.Order {
	return &model.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: order.TransactionUuid,
		PaymentMethod:   model.PaymentMethod(order.PaymentMethod),
		Status:          model.OrderStatus(order.Status),
	}
}

func ModelToRepoModel(status model.OrderStatus) repoModel.OrderStatus {
	return repoModel.OrderStatus(status)
}

func ModelPaymentMethodToRepoModelPaymentMethod(paymentMethod model.PaymentMethod) repoModel.PaymentMethod {
	return repoModel.PaymentMethod(paymentMethod)
}
