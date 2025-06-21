package database

import (
	"entgo.io/ent/dialect"
	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/database/ent"
)

var DbConn *ent.Client

func NewConn() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "")
	exception.Fatal(err, "Failed to connect to database")

	return client
}
