package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-order-service/controllers"
)

func Init(e *echo.Echo) {
	// Routes
	e.POST("/create_order", controllers.CreateOrder)
}
