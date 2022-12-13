package rpcs

import (
	"context"

	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
)

func AuthRPC(userToken string, client user.UserServiceClient) (*user.AuthResponse, error) {

	return client.AuthRPC(context.Background(),
		&user.AuthRequest{
			UserToken: userToken,
		})
}

func UserRPC(userContact string, client user.UserServiceClient) (*user.UserResponse, error) {

	return client.UserRPC(context.Background(),
		&user.UserRequest{
			Contact: userContact,
		})
}
