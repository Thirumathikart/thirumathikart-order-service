package main

import (
	"github.com/thirumathikart/thirumathikart-order-service/config"
	"github.com/thirumathikart/thirumathikart-order-service/registry"
	"github.com/thirumathikart/thirumathikart-order-service/router"

	"github.com/labstack/echo/v4"
)

func main() {

	config.InitApp()

	reg := registry.NewRegistry(config.GetDB())

	server := echo.New()

	router.NewRouter(server, reg.NewAppController())

	server.Logger.Fatal(server.Start(":" + config.ServerPort))

}
