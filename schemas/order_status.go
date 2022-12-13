package schemas

import (
	"database/sql/driver"
)

type OrderStatus string

const (
	BuyerOrdered     OrderStatus = "ORDERED"
	SellerAgreed     OrderStatus = "AGREED"
	DeliveryAssigned OrderStatus = "ASSIGNED"
	DeliveryAccepted OrderStatus = "ACCEPTED"
	DeliveryShipped  OrderStatus = "SHIPPED"
	Delivered        OrderStatus = "DELIVERED"
)

func (os *OrderStatus) Scan(value interface{}) error {
	*os = OrderStatus(value.([]byte))
	return nil
}

func (os OrderStatus) Value() (driver.Value, error) {
	return string(os), nil
}
