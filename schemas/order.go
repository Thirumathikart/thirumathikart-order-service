package schemas

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int    `gorm:"default:0;"`
	AddressID  int    `gorm:"default:0;"`
	Amount     int    `gorm:"default:0;"`
	Status     string `gorm:"default:null;"`
	Stock      int    `gorm:"default:0;"`
}
