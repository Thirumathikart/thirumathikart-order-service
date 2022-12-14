package schemas

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	Order       Order
	OrderID     uint
	Name        string
	CategoryID  uint
	Description string
	Quantity    uint `gorm:"default:0;"`
	Price       uint `gorm:"default:0;"`
}
