package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/services"
	"github.com/thirumathikart/thirumathikart-order-service/utils"
)

type orderController struct {
	service services.OrderService
}

type OrderController interface {
	PlaceOrder(c echo.Context) error
}

func NewOrderController(os services.OrderService) OrderController {
	return &orderController{os}
}

func (os *orderController) PlaceOrder(c echo.Context) error {

	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	request := new(models.CreateOrder)
	if err := c.Bind(request); err != nil {
		return err
	}
	return os.service.AddOrder(c, userDetails, request)
}

func (os *orderController) AcceptOrder(c echo.Context) error {

	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	request := new(models.UpdateOrder)
	if err := c.Bind(request); err != nil {
		return err
	}
	return os.service.AcceptOrder(c, userDetails, request)
}

func (os *orderController) AssignOrder(c echo.Context) error {

	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	request := new(models.UpdateOrder)
	if err := c.Bind(request); err != nil {
		return err
	}
	return os.service.AcceptOrder(c, userDetails, request)
}
