package rpcs

import (
	"context"

	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
)

func AuthRPC(userToken string, client user.UserServiceClient) (*user.GetUserResponse, error) {

	return client.GetUserRPC(context.Background(),
		&user.GetUserRequest{
			UserToken: userToken,
		})
}
