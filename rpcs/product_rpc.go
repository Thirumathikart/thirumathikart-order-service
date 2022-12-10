package rpcs

import (
	"context"

	"github.com/thirumathikart/thirumathikart-order-service/generated/products"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)

func ProductRPC(orderItems []models.CreateOrderItem, client products.ProductServiceClient) (*products.GetProductsResponse, error) {
	var orderProducts = []int32{}
	for _, product := range orderItems {
		orderProducts = append(orderProducts, int32(product.ProductID))
	}
	return client.GetProductsRPC(
		context.Background(),
		&products.GetProductsRequest{
			Products: orderProducts,
		})
}
