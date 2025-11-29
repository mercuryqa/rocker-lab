package order

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/mercuryqa/rocket-lab/order/model"
)

func (r *OrderRepository) GetOrder(ctx context.Context, id string) (*model.GetOrderResponse, bool) {
	// TODO #3 "part_uuids" из базы order_items
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

	var order model.GetOrderResponse

	// TODO &order.PartUuids
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

	jsonData, err := json.Marshal(order)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v\n", err)
		return nil, false
	}

	log.Printf("Order JSON: %s\n", string(jsonData))

	return &order, true
}
