package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HTTPLogger(e *echo.Echo) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} | ${method} ${uri} \t | ${latency_human}\n",
		Output: e.Logger.Output(),
	})
}
