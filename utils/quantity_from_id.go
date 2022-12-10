package utils

import "github.com/thirumathikart/thirumathikart-order-service/models"

func QuantityFromID(items []models.CreateOrderItem) map[int]int {
	quantityFromID := make(map[int]int)
	for _, r := range items {
		quantityFromID[r.ProductID] = r.ProductQuantity
	}
	return quantityFromID
}
