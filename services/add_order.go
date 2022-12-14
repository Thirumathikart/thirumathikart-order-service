package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/generated/notification"
	"github.com/thirumathikart/thirumathikart-order-service/generated/products"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/helpers"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)

func (os *orderService) AddOrder(
	c echo.Context,
	userDetails *user.User,
	request *models.CreateOrder) error {

	notificationChannel := make(chan struct{})

	go func(channel chan struct{}) {
		userRes, err := helpers.GRPCDialler(config.AuthService, "user", request.SellerContact)
		if err != nil {
			log.Panicln(err)
			return
		}
		userResponse := userRes.(*user.UserResponse)
		fcmToken := *userResponse.User.FcmToken
		if fcmToken == "" {
			log.Panicln("Error Occurred")
			return
		}
		notifyRes, err := helpers.GRPCDialler(config.ProductService, "messaging", fcmToken)
		if err != nil {
			log.Panicln(err)
			return
		}
		notifyResponse := notifyRes.(*notification.SingleNotificationResponse)
		if !notifyResponse.IsSuccess {
			msg := fmt.Sprintf("Unable to send notification to fcmToken: %s", fcmToken)
			log.Panicln(msg)
		}
		channel <- struct{}{}
	}(notificationChannel)

	res, err := helpers.GRPCDialler(config.ProductService, "product", request.OrderItems)
	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occurred")
	}

	response := res.(*products.GetProductsResponse)
	err = os.repo.CreateOrder(userDetails, response, request.OrderItems)

	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occurred")
	}

	<-notificationChannel
	return middlewares.Responder(c, http.StatusOK, "Order Placed Successfully")
}
