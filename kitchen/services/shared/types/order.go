package types

type Order struct {
	ID          string  `json:"id"`
	CustomerID  string  `json:"customer_id"`
	ProductID   string  `json:"product_id"`
	Quantity    int32   `json:"quantity"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}

type OrderService interface {
	CreateOrder(customerID, productID string, quantity int32) (*Order, error)
	GetOrder(id string) (*Order, error)
	ListOrders() ([]*Order, error)
}
