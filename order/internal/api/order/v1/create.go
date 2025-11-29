package apiv1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (h *OrderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var req model.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to decode order data", http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateOrder(ctx, &model.OrderRequest{
		UserUuid:  req.UserUuid,
		PartUuids: req.PartUuids,
	})
	if err != nil {
		log.Printf("failed to save order: %v\n", err)
	}

	orderRes := &model.OrderResponse{
		OrderUuid:  resp.OrderUuid,
		TotalPrice: resp.TotalPrice,
	}

	render.JSON(w, r, orderRes)
}
