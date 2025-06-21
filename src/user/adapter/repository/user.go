package repository

import (
	"context"

	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/common/filter"
	"github.com/dapod93/learn-grpc/database/ent"
	"github.com/dapod93/learn-grpc/src/user/domain/abstract"
	"github.com/dapod93/learn-grpc/src/user/domain/entity"
	"github.com/jinzhu/copier"
)

type userRepo struct {
	ctx context.Context
	tx  *ent.Tx
}

func NewUser(ctx context.Context, tx *ent.Tx) abstract.UserRepo {
	return &userRepo{
		ctx: ctx,
		tx:  tx,
	}
}

func (repo *userRepo) GetByFilter(f filter.User) *entity.User {
	user, err := repo.tx.User.Query().Where(f.GetEntPredicates()...).First(repo.ctx)
	exception.General(err, "Failed to run userRepo.GetByFilter")
	return repo.modelToEntity(user)
}

func (repo *userRepo) modelToEntity(model *ent.User) *entity.User {
	if model == nil {
		return nil
	}

	var entity entity.User
	exception.General(copier.Copy(&entity, &model), "Failed to copy model to entity")

	return &entity
}
