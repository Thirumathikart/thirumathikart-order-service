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
	Location          Location
	Status            bool
}

type Location struct {
	gorm.Model
	Lat float64 `gorm:"type:decimal(10,8)"`
	Lng float64 `gorm:"type:decimal(11,8)"`
}
