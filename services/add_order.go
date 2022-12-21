package services

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	//"github.com/thirumathikart/thirumathikart-order-service/generated/notification"
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

	//notificationChannel := make(chan struct{})

	// go func(channel chan struct{}) {
	// 	log.Println("1111111111111111",userDetails)
	// 	userRes, err := helpers.GRPCDialler(config.AuthService, "user", request.SellerContact)
	// 	if err != nil {
	// 		log.Println("444444444444444",err)
	// 		return
	// 	}
	// 	log.Println("22222222222222222",userRes)
	// 	userResponse := userRes.(*user.UserResponse)
	// 	fcmToken := *userResponse.User.FcmToken
	// 	if fcmToken == "" {
	// 		log.Println("555555555555555",err)
	// 		return
	// 	}
	// 	log.Println("333333333333333333",fcmToken)
	// 	notifyRes, err := helpers.GRPCDialler(config.MessagingService, "messaging", fcmToken)
	// 	if err != nil {
	// 		log.Println("66666666666666",err)
	// 		return
	// 	}
	// 	log.Println("333333333333333333",notifyRes)
	// 	notifyResponse := notifyRes.(*notification.SingleNotificationResponse)
	// 	if !notifyResponse.IsSuccess {
	// 		msg := fmt.Sprintf("Unable to send notification to fcmToken: %s", fcmToken)
	// 		log.Println("7777777777777777",msg)
	// 	}
	// 	channel <- struct{}{}
	// }(notificationChannel)
	log.Println("-----------------",request.OrderItems)
	res, err := helpers.GRPCDialler(config.ProductService, "product", request.OrderItems)
	if err != nil {
		log.Println("////////////////",err)
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occurred")
	}
	log.Println("++++++++++++++++",res)
	//{product_id:42  seller_id:3  product_title:"goat"  product_price:1000  category_id:1}

	response := res.(*products.GetProductsResponse)
	log.Println("++++++++++++++++",response)
	err = os.repo.CreateOrder(userDetails, response, request.OrderItems,request.CustomerAddressID)

	if err != nil {
		return middlewares.Responder(c, http.StatusBadRequest, "Error Occurred")
	}

	//<-notificationChannel
	return middlewares.Responder(c, http.StatusOK, "Order Placed Successfully")
}
