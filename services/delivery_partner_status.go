package services

import (
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)

func (os *orderService) DeliveryPartnerStatusService(c echo.Context, request *models.DeliveryPartnerStatus) error {
	return os.repo.UpdateDelvieryPartnerStatus(request.DeliveryPartnerID, request.DeliveryPartnerStatus, request.Lat, request.Lng)
}
