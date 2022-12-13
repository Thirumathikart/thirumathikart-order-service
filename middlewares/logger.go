package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc/grpclog"
)

func HTTPLogger(e *echo.Echo) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} | ${method} ${uri} \t | ${latency_human}\n",
		Output: e.Logger.Output(),
	})
}

var GrpcLog grpclog.LoggerV2

func GrpcLogger() {
	GrpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(GrpcLog)
}
