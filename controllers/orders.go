package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)

// func ListProductsBySeller(sellerID int) []models.Order {
// 	db := config.GetDB()
// 	var products []models.Order
// 	db.Find(&products, "seller_id = ?", sellerID)
// 	return products
// }

// func ListProductsByCategory(categoryID int) []models.Product {
// 	db := config.GetDB()
// 	var products []models.Product
// 	db.Find(&products, "category_id = ?", categoryID)
// 	return products
// }

func CreateOrder(c echo.Context) error {
	requestBody := new(models.CreateOrder)
	if err := c.Bind(requestBody); err != nil {
		return err
	}
	return nil
}
