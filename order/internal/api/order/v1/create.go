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

	//conn, err := grpc.NewClient(
	//	"localhost:50055",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	http.Error(w, "failed to connect to inventory service: "+err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//defer func() {
	//	if cerr := conn.Close(); cerr != nil {
	//		log.Printf("failed to close grpc connection: %v", cerr)
	//	}
	//}()
	//
	//client := inventoryV1.NewInventoryStorageClient(conn)

	//var req model.OrderRequest
	//if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	http.Error(w, "Failed to decode order data", http.StatusBadRequest)
	//	return
	//}

	//resp, err := client.ListParts(ctx, &inventoryV1.GetListPartRequest{
	//	Filter: &inventoryV1.PartsFilter{
	//		Uuids: req.PartUuids,
	//	},
	//})
	//if err != nil {
	//	log.Printf("grpc ListParts call failed: %v", err)
	//	http.Error(w, "inventory service error: "+err.Error(), http.StatusBadGateway)
	//	return
	//}

	render.JSON(w, r, resp)
}
