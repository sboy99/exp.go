package service

import (
	"context"
	"time"

	"google.golang.org/grpc"

	ordersv1 "github.com/sboy99/exp.go/kitchen/gen/go/orders/v1"
	"github.com/sboy99/exp.go/kitchen/services/shared/types"
)

type OrderService struct {
	orderClient ordersv1.OrdersServiceClient
	grpcClient  *grpc.ClientConn
}

func NewOrderService(grpcClient *grpc.ClientConn) *OrderService {
	return &OrderService{
		grpcClient:  grpcClient,
		orderClient: ordersv1.NewOrdersServiceClient(grpcClient),
	}
}

func (s *OrderService) CreateOrder(customerID, productID string, quantity int32) (*types.Order, error) {
	createOrderRequest := &ordersv1.CreateOrderRequest{
		CustomerId: customerID,
		ProductId:  productID,
		Quantity:   quantity,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	orderResponse, err := s.orderClient.CreateOrder(ctx, createOrderRequest)
	if err != nil {
		return nil, err
	}
	return s.GetOrder(orderResponse.OrderId)
}

func (s *OrderService) GetOrder(id string) (*types.Order, error) {
	getOrderRequest := &ordersv1.GetOrderRequest{
		OrderId: id,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	orderResponse, err := s.orderClient.GetOrder(ctx, getOrderRequest)
	if err != nil {
		return nil, err
	}
	order := &types.Order{
		ID:          orderResponse.Order.Id,
		CustomerID:  orderResponse.Order.CustomerId,
		ProductID:   orderResponse.Order.ProductId,
		Quantity:    orderResponse.Order.Quantity,
		TotalAmount: orderResponse.Order.TotalAmount,
		Status:      orderResponse.Order.Status,
	}
	return order, nil
}

func (s *OrderService) ListOrders() ([]*types.Order, error) {
	listOrdersRequest := &ordersv1.ListOrdersRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	orderResponse, err := s.orderClient.ListOrders(ctx, listOrdersRequest)
	if err != nil {
		return nil, err
	}
	var orders []*types.Order
	for _, order := range orderResponse.Orders {
		orders = append(orders, &types.Order{
			ID:          order.Id,
			CustomerID:  order.CustomerId,
			ProductID:   order.ProductId,
			Quantity:    order.Quantity,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
		})
	}
	return orders, nil
}
