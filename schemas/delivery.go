package schemas

import (
	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	Order      Order
	OrderID    uint
	DeliveryID uint
}

type Location struct {
	gorm.Model
	Lat float64 `gorm:"type:decimal(10,8)"`
	Lng float64 `gorm:"type:decimal(11,8)"`
}
