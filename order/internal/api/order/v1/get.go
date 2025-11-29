package apiv1

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *OrderHandler) getOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	id := chi.URLParam(r, "order_uuid")
	if id == "" {
		http.Error(w, "Order UUID parameter is required", http.StatusBadRequest)
		return
	}

	order, ok := h.service.GetOrder(ctx, id)
	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	render.JSON(w, r, order)
}
