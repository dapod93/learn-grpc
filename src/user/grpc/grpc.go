package grpc

import userv1 "github.com/dapod93/learn-grpc/proto/gen/user/v1"

type userGrpcController struct {
	userv1.UnimplementedUserServiceServer
}

func NewUserController() userv1.UserServiceServer {
	return &userGrpcController{}
}
