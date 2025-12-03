package order

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*model.OrderInfo, bool) {
	// Получаю данные из Orders
	builderSelect := sq.
		Select("order_uuid", "user_uuid", "total_price", "transaction_uuid", "payment_method", "status").
		From("orders").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_uuid": id}).Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Printf("Failed to build query: %v\n", err)
		return nil, false
	}

	var order model.OrderInfo

	err = r.poolDb.
		QueryRow(ctx, query, args...).
		Scan(
			&order.OrderUuid,
			&order.UserUuid,
			&order.TotalPrice,
			&order.TransactionUuid,
			&order.PaymentMethod,
			&order.Status,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("no order found")
			return nil, false
		}
		log.Printf("failed to scan order: %v\n", err)
		return nil, false
	}

	// Получаю данные из order_items

	builderSelect = sq.
		Select("part_uuid").
		From("order_items").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_uuid": id})

	query, args, err = builderSelect.ToSql()
	if err != nil {
		log.Printf("Failed to build query: %v\n", err)
		return nil, false
	}

	rows, err := r.poolDb.Query(ctx, query, args...)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return nil, false
	}
	defer rows.Close()

	var partUuids []string
	for rows.Next() {
		var partUuid string
		if err := rows.Scan(&partUuid); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			return nil, false
		}
		partUuids = append(partUuids, partUuid)
	}

	if rows.Err() != nil {
		log.Printf("Rows iteration error: %v\n", rows.Err())
		return nil, false
	}

	// Добавляю items в order

	order.PartUuids = partUuids

	// Логирую

	jsonData, err := json.Marshal(order)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v\n", err)
		return nil, false
	}

	log.Printf("Order JSON: %s\n", string(jsonData))

	// Ответ

	return &order, true
}
