package services

import (
	"fmt"

	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/generated/products"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/utils"
)

type orderService struct{}

type OrderService interface {
	AddOrder(user *user.User,
		productInfo *products.GetProductsResponse,
		req *models.CreateOrder)
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (os *orderService) AddOrder(user *user.User,
	productInfo *products.GetProductsResponse,
	req *models.CreateOrder) {

	quantityFromID := utils.QuantityFromID(req.OrderItems)
	db := config.GetDB()

	order := models.Order{
		CustomerID:  int(user.UserId),
		AddressID:   int(*user.AddressId),
		OrderStatus: models.BuyerOrdered,
	}

	res := db.Create(&order)

	fmt.Println(res)
	orderItems := []models.OrderItem{}
	for _, product := range productInfo.GetProducts() {
		orderItem := models.OrderItem{
			OrderID:  int(order.ID),
			Name:     product.ProductTitle,
			Price:    int(product.ProductPrice),
			Quantity: quantityFromID[int(product.ProductId)],
		}
		orderItems = append(orderItems, orderItem)
	}
	db.Create(&orderItems)
}
