package apiv1

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func (h *OrderHandler) cancelOrder(w http.ResponseWriter, r *http.Request) {
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

	switch order.Status {
	case "PENDING_PAYMENT":
		h.service.CancelOrder(ctx, id, "CANCELED")
		w.WriteHeader(http.StatusNoContent)
		return
	case "PAID":
		http.Error(w, "Order already PAID — cannot change status", http.StatusConflict)
	case "CANCELED":
		http.Error(w, "Order already CANCELED — cannot change status", http.StatusConflict)
	default:
		http.Error(w, "Order not found or invalid status", http.StatusNotFound)
	}
}
