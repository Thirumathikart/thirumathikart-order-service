package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/helpers"
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
	UpdateDeliveryPartnerStatus(c echo.Context) error
	AcceptOrder(c echo.Context) error
	AssignOrder(c echo.Context) error
	ShipOrder(c echo.Context) error
	FetchOrderBySeller(c echo.Context) error
	FetchOrderByDeliveryPartner(c echo.Context) error
	FetchOrderByCustomer(c echo.Context) error
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
	return os.service.AssignOrder(c, userDetails, request)
}

func (os *orderController) ShipOrder(c echo.Context) error {

	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	request := new(models.UpdateOrder)
	if err := c.Bind(request); err != nil {
		return err
	}
	return os.service.ShipOrder(c, userDetails, request)
}

func (os *orderController) UpdateDeliveryPartnerStatus(c echo.Context) error {

	request := new(models.DeliveryPartnerStatus)
	if err := c.Bind(request); err != nil {
		return err
	}
	return os.service.DeliveryPartnerStatusService(c, request)
}

func (os *orderController) FetchOrderBySeller(c echo.Context) error {
	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	return os.service.FetchOrderBySeller(c, userDetails)
}

func (os *orderController) FetchOrderByDeliveryPartner(c echo.Context) error {
	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	// return os.service.FetchOrderByDeliveryPartner(c,userDetails)
	userAddressRequest := &models.UserAddressRequest{
		UserID:    userDetails.UserId,
		AddressID: 2,
	}
	customerRes, err := helpers.GRPCDialler(config.AuthService, "user", userAddressRequest)
	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, err.Error())
	}
	return middlewares.Responder(c, http.StatusOK, customerRes)
}

func (os *orderController) FetchOrderByCustomer(c echo.Context) error {
	userDetails, err := utils.GetUserDetails(c)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	return os.service.FetchOrderByCustomer(c, userDetails)
}
