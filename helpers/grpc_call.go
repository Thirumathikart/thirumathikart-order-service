package helpers

import (
	"fmt"

	"github.com/thirumathikart/thirumathikart-order-service/middlewares"
	"github.com/thirumathikart/thirumathikart-order-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCDialler(url string, rpcType string, request interface{}) (interface{}, error) {
	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		middlewares.WithClientUnaryInterceptor())
	conn, er := grpc.Dial(url, opts...)
	if er != nil {
		fmt.Println("error in dial", er)
	}
	res, err := utils.GRPCSwitch(conn, rpcType, request)
	conn.Close()
	return res, err
}
