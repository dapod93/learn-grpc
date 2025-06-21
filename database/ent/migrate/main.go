//go:build ignore
// +build ignore

package main

import (
	"context"
	"errors"
	"os"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/common/util"
	"github.com/dapod93/learn-grpc/database/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	util.LoadDotEnv(".env")

	if len(os.Args) != 2 {
		exception.Fatal(errors.New("migration new required"), "Migration name is required")
	}

	ctx := context.Background()
	dir, err := atlas.NewLocalDir("database/ent/migrate/migrations")
	exception.Fatal(err, "No such file or directory")

	exception.Fatal(
		migrate.NamedDiff(
			ctx,
			"sqlite://learn-grpc.db?mode=memory&_fk=1",
			os.Args[1],
			[]schema.MigrateOption{
				schema.WithDir(dir),
				schema.WithMigrationMode(schema.ModeReplay),
				schema.WithDialect(dialect.SQLite),
				schema.WithFormatter(atlas.DefaultFormatter),
				schema.WithForeignKeys(false),
			}...),
		"Failed to generate migration file",
	)
}
