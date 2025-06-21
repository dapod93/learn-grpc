package abstract

type UserUOW interface {
	BeginTx()
	Commit()
	Rollback()

	UserRepo() UserRepo
}
