package models

import "github.com/thirumathikart/thirumathikart-order-service/schemas"

type FetchOrder struct {
	Order       schemas.Order 			`json:"order"`
	OrderItem   []schemas.OrderItem 	`json:"items"`
}
