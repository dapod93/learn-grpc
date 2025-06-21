package database

import (
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/dapod93/learn-grpc/common/config"
	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/database/ent"

	_ "github.com/mattn/go-sqlite3"
)

var DbConn *ent.Client

func NewConn() *ent.Client {
	client, err := ent.Open(
		dialect.SQLite,
		fmt.Sprintf("file:%s?_fk=1", config.Server.DatabaseLocation),
	)
	exception.Fatal(err, "Failed to connect to database")

	return client
}
