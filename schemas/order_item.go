package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductId int `gorm:"default:0;"`
	Quantity  int `gorm:"default:0;"`
}
