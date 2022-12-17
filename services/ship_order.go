package services

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/schemas"
)

func (os *orderService) ShipOrder(c echo.Context,
	userDetails *user.User,
	request *models.UpdateOrder) error {

	err := os.repo.UpdateOrderStatus(request.OrderID, schemas.DeliveryShipped)
	if err != nil {
		log.Panicln(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	return middlewares.Responder(c, http.StatusOK, "Order Placed Successfully")

}
