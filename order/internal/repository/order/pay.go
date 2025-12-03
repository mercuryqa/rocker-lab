package order

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *OrderRepository) PayOrder(ctx context.Context, id, status, paymentMethodName, transactionUuid string) bool {
	builderUpdate := sq.Update("orders").
		Set("status", status).
		Set("transaction_uuid", transactionUuid).
		Set("payment_method", paymentMethodName).
		Where(sq.Eq{"order_uuid": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to build query")
		return false
	}

	_, err = r.poolDb.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed update in table orders: %v\n", err)
		return false
	}

	return true
}
