package models

import (
	"database/sql/driver"
)

type OrderStatus string

const (
	BuyerOrdered     OrderStatus = "ordered"
	SellerAgreed     OrderStatus = "agreed"
	DeliveryAssigned OrderStatus = "assigned"
	DeliveryAccepted OrderStatus = "accepted"
	DeliveryShipped  OrderStatus = "shipped"
	Delivered        OrderStatus = "delivered"
)

func (os *OrderStatus) Scan(value interface{}) error {
	*os = OrderStatus(value.([]byte))
	return nil
}

func (os OrderStatus) Value() (driver.Value, error) {
	return string(os), nil
}
