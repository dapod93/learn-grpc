package router

import (
	"google.golang.org/grpc"

	userv1 "github.com/dapod93/learn-grpc/proto/gen/user/v1"
	usergrpc "github.com/dapod93/learn-grpc/src/user/grpc"
)

func RegisterService(server *grpc.Server) {
	userv1.RegisterUserServiceServer(server, usergrpc.NewUserController())
}
