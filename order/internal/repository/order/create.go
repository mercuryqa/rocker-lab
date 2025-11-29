package order

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
	repoModel "github.com/mercuryqa/rocket-lab/order/internal/repository/model"
)

// CreateOrder создает заказ
func (r *OrderRepository) CreateOrder(ctx context.Context, order model.OrderInfo) (string, error) {
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
	orderUuid := uuid.NewString()

	builderInsert := sq.Insert("orders").
		PlaceholderFormat(sq.Dollar).
		Columns("order_uuid", "user_uuid", "total_price", "status", "payment_method").
		Values(orderUuid, order.UserUuid, order.TotalPrice, repoModel.PendingPayment, "").
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

	// вставка items
	builderInsert = sq.Insert("order_items").
		PlaceholderFormat(sq.Dollar).
		Columns("order_uuid", "part_uuid")

	for _, partUuid := range order.PartUuids {
		builderInsert = builderInsert.Values(orderUuidDb, partUuid)
	}

	query, args, err = builderInsert.ToSql()
	if err != nil {
		log.Printf("Ошибка build query: %v\n", err)
	}
	rows, err := r.poolDb.Query(ctx, query, args...)
	if err != nil {
		log.Printf("Ошибка insert в таблицу orders: %v\n", err)
	}
	rows.Close()
	log.Printf("Добавлена запись : %v\n", rows)

	// КОД ДЛЯ ТРАНЗАКЦИИ
	// err = tx.Commit(ctx)
	// if err != nil {
	//	panic(err)
	// }

	return orderUuid, nil
}
