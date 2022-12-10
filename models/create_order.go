package models

type CreateOrderItem struct {
	ProductID       int `json:"product_id"`
	ProductQuantity int `json:"product_quantity"`
}

type CreateOrder struct {
	OrderItems    []CreateOrderItem `json:"order_items"`
	SellerContact string            `json:"seller_contact"`
}
