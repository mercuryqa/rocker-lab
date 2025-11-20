package order

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

// CreateOrder создает заказ
func (s *service) CreateOrder(ctx context.Context, info *model.OrderRequest) (*model.OrderResponse, error) {
	// Получаю запчасти по их id
	parts, err := s.inventoryClient.ListParts(ctx, model.PartsFilter{
		Uuids: info.GetPartUuids(),
	})
	if err != nil {
		return nil, err
	}

	var totalPrice float64

	// Провеока наличия деталей и подсчет суммы
	var existsPartUuids []string
	for _, part := range parts {
		if part.StockQuantity <= 0 {
			continue
		}
		totalPrice += part.Price
		existsPartUuids = append(existsPartUuids, part.UUID)
	}
	if len(existsPartUuids) == 0 {
		log.Printf("No inventory %v\n", err)
		return nil, err
	}

	orderUUID := uuid.New().String()

	order := model.Order{
		OrderUuid:  orderUUID,
		UserUuid:   info.GetUserUuid(),
		PartUuids:  existsPartUuids,
		TotalPrice: totalPrice,
	}

	err = s.orderRepository.CreateOrder(&order)
	if err != nil {
		return nil, err
	}

	return &model.OrderResponse{
		OrderUuid:  order.OrderUuid,
		TotalPrice: order.TotalPrice,
	}, nil
}
