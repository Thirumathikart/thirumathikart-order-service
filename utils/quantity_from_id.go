package utils

import "github.com/thirumathikart/thirumathikart-order-service/models"

func QuantityFromID(items []models.CreateOrderItem) map[uint]uint {
	quantityFromID := make(map[uint]uint)
	for _, r := range items {
		quantityFromID[r.ProductID] = r.ProductQuantity
	}
	return quantityFromID
}
