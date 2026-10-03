package service

import (
	"uuid"

	"github.com/sboy99/exp.go/kitchen/services/shared/types"
)

type OrderService struct {
	orders map[string]*types.Order
}

func NewOrderService() *OrderService {
	return &OrderService{
		orders: make(map[string]*types.Order),
	}
}

func (s *OrderService) CreateOrder(customerID, productID string, quantity int32) (*types.Order, error) {
	order := &types.Order{
		ID:          uuid.New().String(),
		CustomerID:  customerID,
		ProductID:   productID,
		Quantity:    quantity,
		TotalAmount: float64(quantity) * 30, // Assuming each product costs $30
		Status:      "Pending",
	}
	s.orders[order.ID] = order
	return order, nil
}

func (s *OrderService) GetOrder(id string) (*types.Order, error) {
	order, exists := s.orders[id]
	if !exists {
		return nil, nil // or return an error indicating order not found
	}
	return order, nil
}

func (s *OrderService) ListOrders() ([]*types.Order, error) {
	var orderList []*types.Order
	for _, order := range s.orders {
		orderList = append(orderList, order)
	}
	return orderList, nil
}
