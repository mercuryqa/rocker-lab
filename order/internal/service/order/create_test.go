package order

import (
	"reflect"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestCreateSuccess2() {

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

//func (s *ServiceSuite) TestCreateSuccess2() {
//
//	ctx := s.ctx
//
//	var (
//		UserUuid  = "1234"
//		PartUuids = []string{"1", "2", "3"}
//
//		expectedUUID = uuid.New().String()
//
//		req = model.OrderRequest{
//			UserUuid:  UserUuid,
//			PartUuids: PartUuids,
//		}
//
//		// список деталей, которые отдает inventory
//		parts = []model.Part{
//			{UUID: "1", Price: 100, StockQuantity: 1},
//			{UUID: "2", Price: 200, StockQuantity: 5},
//			{UUID: "3", Price: 300, StockQuantity: 2},
//		}
//	)
//
//	// мок inventoryClient
//	s.InventoryClient.
//		On("ListParts", ctx, model.PartsFilter{Uuids: req.PartUuids}).
//		Return(parts, nil)
//
//	//// мок на payment (если CreateOrder вызывает платежи)
//	//s.PaymentClient.
//	//	On("Reserve", s.ctx, UserUuid, PartUuids).
//	//	Return(nil)
//
//	// мок на репозиторий
//	s.OrderRepository.
//		On("CreateOrder", s.ctx, req).
//		Return(expectedUUID, nil)
//
//	resp, err := s.service.CreateOrder(s.ctx, &req)
//	s.Require().NoError(err)
//	s.Require().Equal(expectedUUID, resp)
//
//	// сумма: 100 + 200 + 300 = 600
//	s.Require().Equal(float64(600), resp.TotalPrice)
//	s.Require().NotEmpty(resp.OrderUuid)
//
//	s.InventoryClient.AssertExpectations(s.T())
//	s.OrderRepository.AssertExpectations(s.T())
//
//}
