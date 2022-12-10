package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/generated/products"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/helpers"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/services"
)

type orderController struct {
	order services.OrderService
}

type OrderController interface {
	CreateOrder(c echo.Context) error
}

func NewOrderController(os services.OrderService) OrderController {
	return &orderController{os}
}

func (os *orderController) CreateOrder(c echo.Context) error {

	credentials := c.Get("user").(map[string]interface{})

	var userDetails user.User
	err := mapstructure.Decode(&credentials, &userDetails)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	request := new(models.CreateOrder)
	if err := c.Bind(request); err != nil {
		return err
	}

	res, err := helpers.GRPCDialler(config.ProductService, "product", request.OrderItems)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occurred")
	}
	response := res.(*products.GetProductsResponse)

	os.order.AddOrder(&userDetails, response, request)
	return middlewares.Responder(c, http.StatusOK, response)
}
