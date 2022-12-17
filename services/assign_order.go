package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	priorityQueue "github.com/thirumathikart/thirumathikart-order-service/datastructures"
	"github.com/thirumathikart/thirumathikart-order-service/generated/notification"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/helpers"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/utils"
)

func (os *orderService) AssignOrder(c echo.Context,
	userDetails *user.User,
	request *models.UpdateOrder) error {

	deliveryPartners, err := os.repo.GetDeliveryPartners()
	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "No Delivery Partner is Available now, Try Later")
	}
	if len(deliveryPartners) < 1 {
		return middlewares.Responder(c, http.StatusNoContent, "No Delivery Partner is Available now, Try Later")
	}

	order, err1 := os.repo.GetOrder(request.OrderID)
	if err1 != nil {
		return middlewares.Responder(c, http.StatusNoContent, "No Delivery Partner is Available now, Try Later")
	}
	userResult, err2 := helpers.GRPCDialler(config.AuthService, "user", order.SellerID)
	if err2 != nil {
		return middlewares.Responder(c, http.StatusNoContent, "No Delivery Partner is Available now, Try Later")
	}

	sellerRes := userResult.(*user.UserResponse)
	seller := sellerRes.User.Address
	deliveryPartnersQueue := priorityQueue.New()
	for _, deliveryPartner := range deliveryPartners {
		if deliveryPartner.Status {
			deliveryPartnersQueue.Insert(
				deliveryPartner.ID,
				utils.DistanceCalculator(
					seller.Latitude,
					seller.Longitude,
					deliveryPartner.Latitude,
					deliveryPartner.Longitude,
				))
		}
	}
	deliveryPartner, err3 := deliveryPartnersQueue.Pop()
	if err3 != nil {
		return middlewares.Responder(c, http.StatusNoContent, "No Delivery Partner is available now, Try Later")
	}
	deliveryPartnerID := deliveryPartner.(uint)
	err3 = os.repo.AssignDeliveryPartner(request.OrderID, deliveryPartnerID)
	if err3 != nil {
		return middlewares.Responder(c, http.StatusNoContent, "No Delivery Partner is available now, Try Later")
	}
	userResult, err = helpers.GRPCDialler(config.AuthService, "user", deliveryPartnerID)
	if err != nil {
		log.Panicln(err)
	}
	deliveryRes := userResult.(*user.UserResponse)
	fcmToken := *deliveryRes.User.FcmToken
	if fcmToken == "" {
		log.Panicln("Error Occurred")
	}
	notifyRes, err := helpers.GRPCDialler(config.ProductService, "messaging", fcmToken)
	if err != nil {
		log.Panicln(err)
	}
	notifyResponse := notifyRes.(*notification.SingleNotificationResponse)
	if !notifyResponse.IsSuccess {
		msg := fmt.Sprintf("Unable to send notification to fcmToken: %s", fcmToken)
		log.Panicln(msg)
	}
	return middlewares.Responder(c, http.StatusOK, "Delivery Partner Assigned Successfully")
}
