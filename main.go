package main

import (
	"os"

	"github.com/dapod93/learn-grpc/common/config"
	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/database"

	"buf.build/go/protovalidate"
)

func init() {
	config.Server = config.ServerConfig{
		DatabaseLocation: os.Getenv("DATABASE_LOCATION"),
	}
}

func main() {
	validator, err := protovalidate.New()
	exception.Fatal(err, "Failed to create new validator")

	database.DbConn = database.NewConn()
	defer database.DbConn.Close()
}
