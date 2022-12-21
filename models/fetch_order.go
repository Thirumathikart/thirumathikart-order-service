package models

import (
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/schemas"
)

type FetchOrder struct {
	Order           schemas.Order       `json:"order"`
	OrderItem       []schemas.OrderItem `json:"items"`
	CustomerAddress *user.Address       `json:"customer_address"`
	SellerAddress   *user.Address       `json:"seller_address"`
}
