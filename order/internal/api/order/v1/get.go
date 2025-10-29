package apiV1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *OrderHandler) getOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "order_uuid")
	if id == "" {
		http.Error(w, "Order UUID parameter is required", http.StatusBadRequest)
		return
	}

	order, ok := h.service.GetOrder(id)
	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	render.JSON(w, r, order)
}
