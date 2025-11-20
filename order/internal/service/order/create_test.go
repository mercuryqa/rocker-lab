package order

import (
	"reflect"

	"github.com/stretchr/testify/mock"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (s *ServiceSuite) TestCreateSuccess() {
	ctx := s.ctx

	req := model.OrderRequest{
		UserUuid:  "1234",
		PartUuids: []string{"1", "2", "3"},
	}

	parts := []model.Part{
		{UUID: "1", Price: 100, StockQuantity: 1},
		{UUID: "2", Price: 200, StockQuantity: 5},
		{UUID: "3", Price: 300, StockQuantity: 2},
	}

	// inventory mock
	s.InventoryClient.
		On("ListParts", ctx, model.PartsFilter{Uuids: req.PartUuids}).
		Return(parts, nil)

	// здесь мы захватим UUID и проверим содержимое Order
	var capturedOrderUUID string

	s.OrderRepository.
		On("CreateOrder", mock.MatchedBy(func(o *model.Order) bool {
			// сохранить сгенерированный UUID, чтобы потом сравнить
			capturedOrderUUID = o.OrderUuid

			return o.UserUuid == req.UserUuid &&
				reflect.DeepEqual(o.PartUuids, []string{"1", "2", "3"}) &&
				o.TotalPrice == 600
		})).
		Return(nil)

	resp, err := s.service.CreateOrder(ctx, &req)
	s.Require().NoError(err)

	// проверяем что UUID из репозитория == UUID из ответа
	s.Require().Equal(capturedOrderUUID, resp.OrderUuid)

	// проверяем сумму
	s.Require().Equal(float64(600), resp.TotalPrice)
}
