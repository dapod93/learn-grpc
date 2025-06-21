package grpc

import (
	"context"

	"github.com/dapod93/learn-grpc/common/filter"
	"github.com/dapod93/learn-grpc/common/util"
	"github.com/dapod93/learn-grpc/src/user/adapter/uow"
	"github.com/dapod93/learn-grpc/src/user/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usermv1 "github.com/dapod93/learn-grpc/proto/gen/user/v1/message/v1"
)

func (ctrl *userGrpcController) GetUserByFilter(
	ctx context.Context,
	req *usermv1.GetUserByFilterRequest,
) (resp *usermv1.GetUserByFilterResponse, err error) {
	user := service.GetUserByFilter(
		uow.NewUser(ctx),
		filter.User{ID: util.GetNilPointer(int(util.GetDerefPointer(req.Id))), Email: req.Email},
	)
	if user == nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &usermv1.GetUserByFilterResponse{
		Message: "Get user successful",
		Data: &usermv1.GetUserByFilterResponseData{
			Id:        int64(user.ID),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, nil
}
