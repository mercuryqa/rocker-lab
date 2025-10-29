// internal/api/order_handler.go

package apiV1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
	orderService "github.com/mercuryqa/rocket-lab/order/internal/service"
	"github.com/mercuryqa/rocket-lab/order/model"
	paymentV1 "github.com/mercuryqa/rocket-lab/payment/pkg/proto/payment_v1"
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
		TransactionUuid: "1",
		PaymentMethod:   "UNKNOWN",
		Status:          "PENDING_PAYMENT",
	}

	h.service.CreateOrder(orderSave)

	render.JSON(w, r, orderRes)
}

func toPaymentMethod(s string) paymentV1.PaymentMethod {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "CARD":
		return paymentV1.PaymentMethod_CARD
	case "SBP":
		return paymentV1.PaymentMethod_SBP
	case "CREDIT_CARD":
		return paymentV1.PaymentMethod_CREDIT_CARD
	case "INVESTOR_MONEY":
		return paymentV1.PaymentMethod_INVESTOR_MONEY
	default:
		return paymentV1.PaymentMethod_UNKNOWN
	}
}

func (h *OrderHandler) payOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "order_uuid")
	if id == "" {
		http.Error(w, "Id parameter is required", http.StatusBadRequest)
		return
	}

	order, ok := h.service.GetOrder(id)
	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	if order.Status == "PAID" {
		http.Error(w, "Order already PAID", http.StatusConflict)
		return
	}
	if order.Status == "CANCELED" {
		http.Error(w, "Order canceled - can't be paid", http.StatusConflict)
		return
	}

	var req model.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed decode payment data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "failed to connect to payment service: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Printf("failed to close grpc connection: %v", cerr)
		}
	}()

	client := paymentV1.NewPaymentV1Client(conn)

	method := toPaymentMethod(req.PaymentMethod)
	resp, err := client.PayOrder(ctx, &paymentV1.PayOrderRequest{
		OrderUuid:     order.OrderUuid,
		UserUuid:      order.UserUuid,
		PaymentMethod: method,
	})
	if err != nil {
		log.Printf("grpc call failed: %v", err)
		http.Error(w, "payment service error: "+err.Error(), http.StatusBadGateway)
		return
	}

	h.service.PayOrder(id, "PAID")
	paymentMethodName := paymentV1.PaymentMethod_name[int32(method)]
	h.service.UpdateOrder(id, paymentMethodName, resp.TransactionUuid)

	render.JSON(w, r, map[string]interface{}{
		"transaction_uuid": resp.TransactionUuid,
	})
}

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

func (h *OrderHandler) cancelOrder(w http.ResponseWriter, r *http.Request) {
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

	switch order.Status {
	case "PENDING_PAYMENT":
		h.service.CanсelOrder(id, "CANCELED")
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
