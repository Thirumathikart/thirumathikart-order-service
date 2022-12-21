package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/generated/notification"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/helpers"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/schemas"
)

func (os *orderService) AcceptOrder(c echo.Context,
	userDetails *user.User,
	request *models.UpdateOrder) error {

	notificationChannel := make(chan struct{})
	clientID, err := os.repo.FindCustomer(request.OrderID)

	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	go func(notiChan chan struct{}, clientID uint) {
		userRes, err := helpers.GRPCDialler(config.AuthService, "user", clientID)
		if err != nil {
			log.Panicln(err)
			notiChan <- struct{}{}
			return
		}
		userResponse := userRes.(*user.UserResponse)
		fcmToken := *userResponse.User.FcmToken
		if fcmToken == "" {
			log.Panicln("Error Occurred")
			notiChan <- struct{}{}
			return
		}
		notifyRes, err := helpers.GRPCDialler(config.MessagingService, "messaging", fcmToken)
		if err != nil {
			log.Panicln(err)
			notiChan <- struct{}{}
			return
		}
		notifyResponse := notifyRes.(*notification.SingleNotificationResponse)
		if !notifyResponse.IsSuccess {
			msg := fmt.Sprintf("Unable to send notification to fcmToken: %s", fcmToken)
			log.Panicln(msg)
		}
		notiChan <- struct{}{}
	}(notificationChannel, clientID)

	err = os.repo.UpdateOrderStatus(request.OrderID, schemas.SellerAgreed)
	if err != nil {
		log.Panicln(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	<-notificationChannel
	return middlewares.Responder(c, http.StatusOK, "Order Placed Successfully")

}
