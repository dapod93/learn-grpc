package service

import (
	"github.com/dapod93/learn-grpc/common/filter"
	"github.com/dapod93/learn-grpc/src/user/domain/abstract"
	"github.com/dapod93/learn-grpc/src/user/domain/entity"
)

func GetUserByFilter(uow abstract.UserUOW, f filter.User) *entity.User {
	// To close the transactions, better to use one transactions per service call
	defer uow.Rollback()
	return uow.UserRepo().GetByFilter(f)
}
