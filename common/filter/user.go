package filter

import (
	"github.com/dapod93/learn-grpc/database/ent/predicate"
	"github.com/dapod93/learn-grpc/database/ent/user"
)

type User struct {
	ID    *int
	Email *string
}

func (uf *User) GetEntPredicates() (predicates []predicate.User) {
	if uf.ID != nil {
		predicates = append(predicates, user.IDEQ(*uf.ID))
	}

	if uf.Email != nil {
		predicates = append(predicates, user.EmailEQ(*uf.Email))
	}

	return predicates
}
