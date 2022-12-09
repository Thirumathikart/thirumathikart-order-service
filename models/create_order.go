package models

type CreateOrderItem struct {
	Product_Title    string `json:"product_title"`
	Product_Quantity string `json:"product_quantity"`
}

type CreateOrder struct {
	Order_Items    []CreateOrderItem `json:"order_items"`
	Seller_Contact string            `json:"seller_contact"`
}
