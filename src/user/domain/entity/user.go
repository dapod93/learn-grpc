package entity

import "time"

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	CreatedAt time.Time
}
