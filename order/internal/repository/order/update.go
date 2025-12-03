package order

// import (
//	"github.com/mercuryqa/rocket-lab/order/internal/converter"
//	"github.com/mercuryqa/rocket-lab/order/internal/model"
// )
//
// func (r *OrderRepository) UpdateOrder(id string, paymentMethod model.PaymentMethod, transactionUuid string) bool {
//	r.mu.Lock()
//	defer r.mu.Unlock()
//
//	order, ok := r.orders[id]
//	if !ok {
//		return false
//	}
//
//	order.TransactionUuid = transactionUuid
//	order.PaymentMethod = converter.ToRepoModelPaymentMethodModelPaymentMethod(paymentMethod)
//
//	return true
// }
