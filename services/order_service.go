package services

import (
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/repositories"
)

type orderService struct {
	repo repositories.OrderRepository
}

type OrderService interface {
	AddOrder(
		c echo.Context,
		userDetails *user.User,
		request *models.CreateOrder) error

	AcceptOrder(
		c echo.Context,
		userDetails *user.User,
		request *models.UpdateOrder) error

	AssignOrder(
		c echo.Context,
		userDetails *user.User,
		request *models.UpdateOrder) error

	ShipOrder(
		c echo.Context,
		userDetails *user.User,
		request *models.UpdateOrder) error
	DeliveryPartnerStatusService(
		c echo.Context,
		request *models.DeliveryPartnerStatus) error

	FetchOrderBySeller(
		c echo.Context,
		userDetails *user.User) error

	FetchOrderByCustomer(
		c echo.Context,
		userDetails *user.User) error

	FetchOrderByDeliveryPartner(
		c echo.Context,
		userDetails *user.User) error
		
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo: repo}
}
