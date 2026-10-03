package handler

import (
	"context"
	"log"

	ordersv1 "github.com/sboy99/exp.go/kitchen/gen/go/orders/v1"
	"github.com/sboy99/exp.go/kitchen/services/shared/types"
	"google.golang.org/grpc"
)

type OrderGRPCHandler struct {
	service types.OrderService
	ordersv1.UnimplementedOrdersServiceServer
}

func NewOrderGRPCHandler(grpcServer *grpc.Server, service types.OrderService) {
	handler := &OrderGRPCHandler{
		service: service,
	}
	// Register the handler with the gRPC server
	ordersv1.RegisterOrdersServiceServer(grpcServer, handler)
}

func (h *OrderGRPCHandler) CreateOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	createdOrder, err := h.service.CreateOrder(req.CustomerId, req.ProductId, req.Quantity)
	if err != nil {
		return nil, err
	}
	log.Printf("Created order: %+v", createdOrder)
	return &ordersv1.CreateOrderResponse{
		Status:  "success",
		OrderId: createdOrder.ID,
	}, nil
}

func (h *OrderGRPCHandler) GetOrder(ctx context.Context, req *ordersv1.GetOrderRequest) (*ordersv1.GetOrderResponse, error) {
	order, err := h.service.GetOrder(req.OrderId)
	if err != nil {
		return nil, err
	}
	log.Printf("Retrieved order: %+v", order)
	return &ordersv1.GetOrderResponse{
		Order: &ordersv1.Order{
			Id:          order.ID,
			CustomerId:  order.CustomerID,
			ProductId:   order.ProductID,
			Quantity:    order.Quantity,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
		},
	}, nil
}

func (h *OrderGRPCHandler) ListOrders(ctx context.Context, req *ordersv1.ListOrdersRequest) (*ordersv1.ListOrdersResponse, error) {
	orders, err := h.service.ListOrders()
	if err != nil {
		return nil, err
	}
	log.Printf("Retrieved orders: %+v", orders)

	var orderList []*ordersv1.Order
	for _, order := range orders {
		orderList = append(orderList, &ordersv1.Order{
			Id:          order.ID,
			CustomerId:  order.CustomerID,
			ProductId:   order.ProductID,
			Quantity:    order.Quantity,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
		})
	}

	return &ordersv1.ListOrdersResponse{
		Orders: orderList,
	}, nil
}
