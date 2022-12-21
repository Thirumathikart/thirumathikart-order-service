package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/helpers"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)

func (os *orderService) FetchOrderBySeller(c echo.Context,
	userDetails *user.User) error {

	orders, err := os.repo.FetchOrderBySeller(uint(userDetails.UserId))

	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
	}

	var response []models.FetchOrder

	for _, order := range orders {
		items, err := os.repo.FetchOrderItemsByOrder(order.ID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		customerRes, err := helpers.GRPCDialler(config.AuthService, "user", order.CustomerID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		sellerRes, err := helpers.GRPCDialler(config.AuthService, "user", order.SellerID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		fmt.Print((customerRes.(*user.UserResponse)).User, (sellerRes.(*user.UserResponse)).User)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		CustomerAddress := customerRes.(*user.UserResponse).User.Address
		SellerAddress := sellerRes.(*user.UserResponse).User.Address
		response = append(response, models.FetchOrder{Order: order, OrderItem: items, CustomerAddress: CustomerAddress, SellerAddress: SellerAddress})
	}

	//TODO: fetch address and user details if required
	return middlewares.Responder(c, http.StatusOK, response)
}

func (os *orderService) FetchOrderByDeliveryPartner(c echo.Context,
	userDetails *user.User) error {

	orders, err := os.repo.FetchOrderByDeliveryPartner(uint(userDetails.UserId))

	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
	}
	var response []models.FetchOrder

	for _, order := range orders {
		items, err := os.repo.FetchOrderItemsByOrder(order.ID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		response = append(response, models.FetchOrder{Order: order, OrderItem: items})
	}

	//TODO: fetch address and user details if required
	return middlewares.Responder(c, http.StatusOK, response)
}

func (os *orderService) FetchOrderByCustomer(c echo.Context,
	userDetails *user.User) error {

	orders, err := os.repo.FetchOrderByCustomer(uint(userDetails.UserId))

	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
	}
	var response []models.FetchOrder

	for _, order := range orders {
		items, err := os.repo.FetchOrderItemsByOrder(order.ID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		response = append(response, models.FetchOrder{Order: order, OrderItem: items})
	}

	//TODO: fetch address and user details if required
	return middlewares.Responder(c, http.StatusOK, response)
}
