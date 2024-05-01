package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx",
		"postgres://postgres:123@localhost/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m, err := migrate.New(
		"file://database/migrations",
		"postgres://username:password@localhost/dbname",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	fmt.Println("Migration completed")
}
