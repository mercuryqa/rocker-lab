package apiv1

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (a *APISuite) TestAPI_Get() {
	id := "order-123"

	expectedOrder := &model.Order{
		OrderUuid:  id,
		UserUuid:   "user-1",
		PartUuids:  []string{"p1", "p2"},
		TotalPrice: 150.0,
	}

	a.orderService.
		On("GetOrder", id).
		Return(expectedOrder, true)

	// создаем запрос
	r := httptest.NewRequest("GET", "/"+id, nil)

	// подсовываем chi route params
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("order_uuid", id)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, routeCtx))

	w := httptest.NewRecorder()

	a.api.getOrder(w, r)

	a.Require().Equal(http.StatusOK, w.Code)

	var got model.Order
	err := json.Unmarshal(w.Body.Bytes(), &got)
	a.Require().NoError(err)

	a.Require().Equal(*expectedOrder, got)

	a.orderService.AssertExpectations(a.T())
}

func (a *APISuite) TestGetOrderNotFound() {
	id := "missing-order"

	a.orderService.
		On("GetOrder", id).
		Return((*model.Order)(nil), false)

	r := httptest.NewRequest("GET", "/"+id, nil)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("order_uuid", id)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, routeCtx))

	w := httptest.NewRecorder()

	a.api.getOrder(w, r)

	a.Require().Equal(http.StatusNotFound, w.Code)

	// http.Error добавляет \n в конец
	a.Require().Equal("Order not found\n", w.Body.String())

	a.orderService.AssertExpectations(a.T())
}
