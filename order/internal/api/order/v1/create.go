package apiV1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
	"github.com/mercuryqa/rocket-lab/order/model"
)

func (h *OrderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50055",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "failed to connect to inventory service: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Printf("failed to close grpc connection: %v", cerr)
		}
	}()

	client := inventoryV1.NewInventoryStorageClient(conn)

	var req model.OrderRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to decode order data", http.StatusBadRequest)
		return
	}

	totalPrice := 0.0
	for _, partUUID := range req.PartUuids {
		resp, err := client.GetPart(ctx, &inventoryV1.GetPartRequest{InventoryUuid: partUUID})
		if err != nil {
			log.Printf("grpc call failed: %v", err)
			http.Error(w, "inventory service error: "+err.Error(), http.StatusBadGateway)
			return
		}
		totalPrice += resp.Part.Price
	}

	orderUUID := uuid.New().String()

	orderRes := &model.OrderResponse{
		OrderUuid:  orderUUID,
		TotalPrice: totalPrice,
	}

	orderSave := &model.GetOrderResponse{
		OrderUuid:       orderUUID,
		UserUuid:        req.UserUuid,
		PartUuids:       req.PartUuids,
		TotalPrice:      totalPrice,
		TransactionUuid: "",
		PaymentMethod:   "UNKNOWN",
		Status:          "PENDING_PAYMENT",
	}

	h.service.CreateOrder(orderSave)

	render.JSON(w, r, orderRes)
}
