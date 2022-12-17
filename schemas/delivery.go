package schemas

import (
	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	Order             Order
	OrderID           uint
	DeliveryPartner   DeliveryPartner
	DeliveryPartnerID uint
}

type DeliveryPartner struct {
	gorm.Model
	DeliveryPartnerID uint
	Latitude          float64 `gorm:"type:decimal(10,8)"`
	Longitude         float64 `gorm:"type:decimal(11,8)"`
	Status            bool
}
