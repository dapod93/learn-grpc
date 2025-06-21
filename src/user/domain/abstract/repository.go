package abstract

import (
	"github.com/dapod93/learn-grpc/common/filter"
	"github.com/dapod93/learn-grpc/src/user/domain/entity"
)

type UserRepo interface {
	GetByFilter(f filter.User) *entity.User
}
