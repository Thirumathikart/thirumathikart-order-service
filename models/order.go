package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID  int         `gorm:"default:0;"`
	AddressID   int         `gorm:"default:0;"`
	SellerID    int         `gorm:"default:0;"`
	OrderStatus OrderStatus `sql:"type:order_status"`
}

func (Order) TableName() string {
	return "order"
}
