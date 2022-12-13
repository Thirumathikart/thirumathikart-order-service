package schemas

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID     uint
	Order       Order
	Name        string
	CategoryID  uint
	Description string
	Quantity    uint `gorm:"default:0;"`
	Price       uint `gorm:"default:0;"`
}
