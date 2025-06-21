package uow

import (
	"context"

	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/common/util"
	"github.com/dapod93/learn-grpc/database"
	"github.com/dapod93/learn-grpc/database/ent"
	"github.com/dapod93/learn-grpc/src/user/adapter/repository"
	"github.com/dapod93/learn-grpc/src/user/domain/abstract"
)

type userUow struct {
	ctx context.Context
	db  *ent.Client
	tx  *ent.Tx

	isTxClosed bool
}

func NewUser(ctx context.Context) abstract.UserUOW {
	uow := &userUow{
		ctx: ctx,
		db:  database.DbConn,
	}
	uow.BeginTx()

	return uow
}

// UserRepo implements abstract.UserUOW.
func (uow *userUow) UserRepo() abstract.UserRepo {
	if uow.isTxClosed {
		uow.BeginTx()
	}

	return repository.NewUser(uow.ctx, uow.tx)
}

// BeginTx implements abstract.UserUOW.
func (uow *userUow) BeginTx() {
	uow.tx = util.MustSet(uow.db.Tx(uow.ctx))
}

// Commit implements abstract.UserUOW.
func (uow *userUow) Commit() {
	uow.isTxClosed = true
	exception.General(uow.tx.Commit(), "Failed to commit transaction")
}

// Rollback implements abstract.UserUOW.
func (uow *userUow) Rollback() {
	uow.isTxClosed = true
	exception.General(uow.tx.Rollback(), "Failed to rollback transaction")
}

// RollbackRecover implements abstract.UserUOW.
func (uow *userUow) RollbackRecover() {
	if r := recover(); r != nil {
		uow.Rollback()
		panic(r)
	}
}
