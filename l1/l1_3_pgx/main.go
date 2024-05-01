package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@localhost/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(
		context.Background(), `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	age INTEGER NOT NULL
	)
	`,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Exec(context.Background(), "INSERT INTO users (name, age) VALUES($1, $2)", "Nik", 21)
	if err != nil {
		log.Fatal("1" + err.Error())
	}
	row := conn.QueryRow(context.Background(), "select * from users where id=$1", 1)
	var name string
	var id, age int
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
