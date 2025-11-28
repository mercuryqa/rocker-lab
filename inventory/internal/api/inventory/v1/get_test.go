package apiv1

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercuryqa/rocket-lab/inventory/internal/converter"
	"github.com/mercuryqa/rocket-lab/inventory/internal/model"
	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

func (a *APISuite) TestGetPartSuccess() {
	id := "order-123"

	partModel := model.Part{
		UUID:          "123",
		Name:          "name",
		Description:   "description",
		Price:         100.00,
		StockQuantity: 1,
	}

	// Мокаем сервисный слой
	a.inventoryService.
		On("GetPart", a.ctx, id).
		Return(partModel, nil)

	req := &inventoryV1.GetPartRequest{
		InventoryUuid: id,
	}

	resp, err := a.api.GetPart(a.ctx, req)

	a.Require().NoError(err)
	a.Require().NotNil(resp)

	expectedProto := converter.ToProtoPart(&partModel)
	a.Require().Equal(expectedProto, resp.Part)

	a.inventoryService.AssertExpectations(a.T())
}

func (a *APISuite) TestGetPartNotFound() {
	id := "missing-part"

	// Сервис возвращает ErrNotFound
	a.inventoryService.
		On("GetPart", mock.Anything, id).
		Return(model.Part{}, model.ErrNotFound)

	req := &inventoryV1.GetPartRequest{
		InventoryUuid: id,
	}

	resp, err := a.api.GetPart(a.ctx, req)

	a.Require().Nil(resp)
	a.Require().Error(err)

	st, ok := status.FromError(err)
	a.Require().True(ok)
	a.Require().Equal(codes.NotFound, st.Code())
	a.Require().Equal("part not found", st.Message())

	a.inventoryService.AssertExpectations(a.T())
}
