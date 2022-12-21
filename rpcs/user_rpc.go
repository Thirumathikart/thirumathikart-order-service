package rpcs

import (
	"context"

	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
	"github.com/thirumathikart/thirumathikart-order-service/models"
)

func AuthRPC(userToken string, client user.UserServiceClient) (*user.AuthResponse, error) {

	return client.AuthRPC(context.Background(),
		&user.AuthRequest{
			UserToken: userToken,
		})
}

func UserRPC(userAddressRequest *models.UserAddressRequest, client user.UserServiceClient) (*user.UserResponse, error) {

	return client.UserRPC(context.Background(),
		&user.UserRequest{
			UserID:    uint32(userAddressRequest.UserID),
			AddressID: uint32(userAddressRequest.AddressID),
		})
}
