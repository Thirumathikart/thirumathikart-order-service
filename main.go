package main

import (
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitConfig()

	config.ConnectDB()
	config.MigrateDB()

	server := echo.New()
	middlewares.InitLogger(server)
	server.Use(middleware.CORS())

	routes.Init(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
