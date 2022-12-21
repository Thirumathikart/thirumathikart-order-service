package schemas

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID        uint        `gorm:"default:0;"`
	CustomerAddressID uint        `gorm:"default:0;"`
	SellerID          uint        `gorm:"default:0;"`
	SellerAddressID   uint        `gorm:"default:0;"`
	DeliveryPartnerID uint        `gorm:"default:0;"`
	OrderStatus       OrderStatus `gorm:"type:order_status"`
}

func (Order) TableName() string {
	return "order"
}
