package models

type CreateOrderItem struct {
	ProductID       uint `json:"product_id"`
	ProductQuantity uint `json:"product_quantity"`
}

type CreateOrder struct {
	OrderItems    []CreateOrderItem `json:"order_items"`
	SellerContact string            `json:"seller_contact"`
	CustomerAddressID uint       	`json:"address_id"`
}
