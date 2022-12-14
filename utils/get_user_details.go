package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/thirumathikart/thirumathikart-order-service/generated/user"
)

func GetUserDetails(c echo.Context) (*user.User, error) {
	credentials := c.Get("user").(map[string]interface{})

	var userDetails user.User
	err := mapstructure.Decode(&credentials, &userDetails)
	return &userDetails, err
}
