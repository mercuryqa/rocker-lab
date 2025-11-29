package order

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	repoModel "github.com/mercuryqa/rocket-lab/order/internal/repository/model"
	"github.com/mercuryqa/rocket-lab/order/model"
)

// CreateOrder создает заказ
func (r *OrderRepository) CreateOrder(ctx context.Context, order *model.GetOrderResponse) (string, error) {
	// КОД ДЛЯ ТРАНЗАКЦИИ
	// tx, err := r.poolDb.BeginTx(ctx, pgx.TxOptions{})
	// if err != nil {
	//	panic(err)
	// }
	// defer func() {
	//	err = tx.Rollback(ctx)
	//	if err != nil {
	//		log.Printf("Ошибка отмены tr: %v\n", err)
	//	}
	// }()

	// log.Printf("IDS %v\n", ids)
	//
	// if !r.CheckItems(ctx, ids) {
	//	log.Printf("failed find items")
	//	return "", nil
	//}

	builderInsert := sq.Insert("orders").
		PlaceholderFormat(sq.Dollar).
		Columns("order_uuid", "user_uuid", "total_price", "status", "payment_method").
		Values(order.OrderUuid, order.UserUuid, order.TotalPrice, repoModel.PendingPayment, "").
		Suffix("RETURNING order_uuid")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("Ошибка build query: %v\n", err)
		return "", err
	}

	var orderUuidDb string
	err = r.poolDb.QueryRow(ctx, query, args...).Scan(&orderUuidDb)
	if err != nil {
		log.Printf("Ошибка insert в таблицу orders: %v\n", err)
		return "", err
	}

	// КОД ДЛЯ ТРАНЗАКЦИИ
	// err = tx.Commit(ctx)
	// if err != nil {
	//	panic(err)
	// }

	return order.OrderUuid, nil
}
