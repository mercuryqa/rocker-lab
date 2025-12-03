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
		http.Error(w, "failed decode order request", http.StatusBadRequest)
	}

	var resp *model.OrderResponse
	resp, err := h.service.CreateOrder(ctx, &req)
	if err != nil {
		log.Printf("failed call CreateOrder")
	}

	render.JSON(w, r, resp)
}
