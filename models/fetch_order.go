package models

import (
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/schemas"
)

type FetchOrder struct {
	Order       schemas.Order 			`json:"order"`
	OrderItem   []schemas.OrderItem 	`json:"items"`
	Customer	*user.User       	`json:"customer"`
	Seller	   	*user.User       	`json:"seller"`
}

type UserAddressRequest struct {
	UserID    uint32 `json:"user_id"`
	AddressID uint32 `json:"address_id"`
}
