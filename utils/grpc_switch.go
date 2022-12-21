package utils

import (
	"github.com/thirumathikart/thirumathikart-order-service/generated/notification"
	"github.com/thirumathikart/thirumathikart-order-service/generated/products"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/models"
	"github.com/thirumathikart/thirumathikart-order-service/rpcs"
	"google.golang.org/grpc"
)

func GRPCSwitch(conn *grpc.ClientConn, rpcType string, request interface{}) (interface{}, error) {
	var response interface{}
	var err error
	switch rpcType {
	case "product":
		req := request.([]models.CreateOrderItem)
		prodClient := products.NewProductServiceClient(conn)
		response, err = rpcs.ProductRPC(req, prodClient)
	case "auth":
		req := request.(string)
		userClient := user.NewUserServiceClient(conn)
		response, err = rpcs.AuthRPC(req, userClient)
	case "user":
		req := request.(*models.UserAddressRequest)
		userClient := user.NewUserServiceClient(conn)
		response, err = rpcs.UserRPC(req, userClient)
	default:
		req := request.(*notification.SingleNotificationRequest)
		notifyClient := notification.NewNotificationServiceClient(conn)
		response, err = rpcs.NotificationRPC(req, notifyClient)
	}
	return response, err
}
