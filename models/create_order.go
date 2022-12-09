package models

type CreateOrderItem struct {
	ProductTitle    string `json:"product_title"`
	ProductQuantity string `json:"product_quantity"`
}

type CreateOrder struct {
	OrderItems    []CreateOrderItem `json:"order_items"`
	SellerContact string            `json:"seller_contact"`
}
