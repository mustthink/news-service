package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	var storageUri, migrationsPath, migrationsTable string

	flag.StringVar(&storageUri, "storage-path", "", "storage uri for migrations")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "schema_migrations", "name of migrations")
	flag.Parse()

	if isEmpty(storageUri, migrationsPath) {
		panic("require flags is empty")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("%s?sslmode=disable&x-migrations-table=%s", storageUri, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("no migrations to apply")
			return
		}

		panic(err)
	}
	println("successfully migrated")
}

func isEmpty(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return true
		}
	}
	return false
}
