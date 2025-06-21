package abstract

type UserUOW interface {
	BeginTx()
	Commit()
	Rollback()
	RollbackRecover()

	UserRepo() UserRepo
}
