package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID     int
	Order       Order
	Name        string
	CategoryID  int
	Description string
	Quantity    int `gorm:"default:0;"`
	Price       int `gorm:"default:0;"`
}
