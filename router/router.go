package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thirumathikart/thirumathikart-order-service/controllers"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
)

func NewRouter(e *echo.Echo, c controllers.OrderController) {
	// Router Middlewares
	e.Use(middleware.CORS())
	e.Use(middlewares.HTTPLogger(e))
	e.Use(middlewares.Authenticator)

	// Router Routes
	e.POST("/create_order", c.PlaceOrder)
	e.POST("/update-deliver-partner-status", c.UpdateDeliveryPartnerStatus)
}
