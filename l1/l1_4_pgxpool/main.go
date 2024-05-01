package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Создание пула соединений
	pool, err := pgxpool.Connect(context.Background(), "postgres://postgres:123@localhost/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	// Получение соединения из пула
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()
	// Использование соединения
	row := conn.QueryRow(context.Background(), "select * from users where id=$1", 1)
	var name string
	var id, age int
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
