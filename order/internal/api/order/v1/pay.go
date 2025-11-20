package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/mercuryqa/rocket-lab/order/internal/model"
)

func (h *OrderHandler) payOrder(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "order_uuid")
	if uuid == "" {
		http.Error(w, "Id parameter is required", http.StatusBadRequest)
		return
	}

	var req model.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed decode payment data", http.StatusBadRequest)
		return
	}

	transactionUuid := h.service.PayOrder(uuid, req.PaymentMethod)

	render.JSON(w, r, map[string]interface{}{
		"transaction_uuid": transactionUuid,
	})
}
