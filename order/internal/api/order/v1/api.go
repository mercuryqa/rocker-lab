// internal/api/order_handler.go

package apiV1

import (
	"github.com/go-chi/chi/v5"

	orderService "github.com/mercuryqa/rocket-lab/order/internal/service"
)

type OrderHandler struct {
	service orderService.OrderService
}

func NewOrderHandler(service orderService.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/orders", func(r chi.Router) {
		r.Post("/", h.createOrder)
		r.Post("/{order_uuid}/pay", h.payOrder)
		r.Get("/{order_uuid}", h.getOrder)
		r.Post("/{order_uuid}/cancel", h.cancelOrder)
	})
}
