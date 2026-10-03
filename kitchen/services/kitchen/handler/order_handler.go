package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sboy99/exp.go/kitchen/services/kitchen/service"
	"google.golang.org/grpc"
)

type OrderHttpHandler struct {
	service *service.OrderService
}

func NewOrderHttpHandler(grpcClient *grpc.ClientConn) *OrderHttpHandler {
	return &OrderHttpHandler{
		service: service.NewOrderService(grpcClient),
	}
}

func (h *OrderHttpHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received HTTP request to create order")
	order, err := h.service.CreateOrder("customer1", "product1", 2)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}
	orderResponse, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal order response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(orderResponse)
}

func (h *OrderHttpHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received HTTP request to get order")
	order, err := h.service.GetOrder("order1")
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}
	orderResponse, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal order response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(orderResponse)
}

func (h *OrderHttpHandler) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received HTTP request to list orders")
	orders, err := h.service.ListOrders()
	if err != nil {
		http.Error(w, "Failed to list orders", http.StatusInternalServerError)
		return
	}
	orderResponse, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Failed to marshal order response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(orderResponse)
}
