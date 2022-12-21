package services

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)


func (os *orderService) FetchOrderBySeller (c echo.Context,
	userDetails *user.User) error {
	
	orders, err :=os.repo.FetchOrderBySeller(uint(userDetails.UserId))

	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
	}

	var response []models.FetchOrder

	for _, order := range orders {
		items, err := os.repo.FetchOrderItemsByOrder(order.ID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		response = append(response, models.FetchOrder{Order: order,OrderItem: items})
	}


	//TODO: fetch address and user details if required
	return middlewares.Responder(c, http.StatusOK,  response)
}

func (os *orderService) FetchOrderByDeliveryPartner (c echo.Context,
	userDetails *user.User) error {
	
	orders, err :=os.repo.FetchOrderByDeliveryPartner(uint(userDetails.UserId))

	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
	}
	var response []models.FetchOrder

	for _, order := range orders {
		items, err := os.repo.FetchOrderItemsByOrder(order.ID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		response = append(response, models.FetchOrder{Order: order,OrderItem: items})
	}


	//TODO: fetch address and user details if required
	return middlewares.Responder(c, http.StatusOK,  response)
}

func (os *orderService) FetchOrderByCustomer (c echo.Context,
	userDetails *user.User) error {
	log.Print(userDetails)
	orders, err :=os.repo.FetchOrderByCustomer(uint(userDetails.UserId))
	log.Print(orders)
	if err != nil {
		return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
	}
	var response []models.FetchOrder

	for _, order := range orders{
		log.Print(order)
		items, err := os.repo.FetchOrderItemsByOrder(order.ID)
		if err != nil {
			return middlewares.Responder(c, http.StatusNoContent, "Unable to Fetch Order")
		}
		response = append(response, models.FetchOrder{Order: order,OrderItem: items})
	}

	//TODO: fetch address and user details if required
	return middlewares.Responder(c, http.StatusOK,  response)
}

