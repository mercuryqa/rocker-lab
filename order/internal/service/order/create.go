package order

import (
	"context"
	"errors"
	"slices"
	"strconv"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

// CreateOrder создает заказ
func (s *service) CreateOrder(ctx context.Context, info *model.OrderRequest) (*model.OrderResponse, error) {
	var totalPrice float64

	parts, err := s.inventoryClient.ListParts(ctx, model.PartsFilter{
		Uuids: info.PartUuids,
	})
	if err != nil {
		return nil, err
	}

	var existsPartUuids []string
	for partUuid, part := range parts {
		if part.StockQuantity <= 0 {
			continue
		}
		totalPrice += part.Price
		partUuidstr := strconv.Itoa(partUuid)
		existsPartUuids = append(existsPartUuids, partUuidstr)
	}
	if len(existsPartUuids) == 0 {
		return nil, errors.New("ошибка: запчасти недоступны")
	}

	slices.Sort(existsPartUuids)

	order := model.OrderInfo{
		UserUuid:   info.UserUuid,
		PartUuids:  existsPartUuids,
		TotalPrice: totalPrice,
	}

	orderUuid, err := s.orderRepository.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &model.OrderResponse{
		OrderUuid:  orderUuid,
		TotalPrice: order.TotalPrice,
	}, nil
}
