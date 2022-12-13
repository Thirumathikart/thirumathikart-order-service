package schemas

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID  uint        `gorm:"default:0;"`
	AddressID   uint        `gorm:"default:0;"`
	SellerID    uint        `gorm:"default:0;"`
	OrderStatus OrderStatus `sql:"type:order_status"`
}

func (Order) TableName() string {
	return "order"
}
