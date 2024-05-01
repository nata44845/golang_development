package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	db, err := sqlx.Connect("pgx", "postgres://postgres:123@localhost/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы
	db.MustExec(`CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	age INTEGER NOT NULL
	)`)
	// Добавление пользователя
	db.MustExec("INSERT INTO users (name, age) VALUES ($1, $2)", "John Doe", 25)
	// Получение списка пользователей
	users := []User{}
	err = db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name,
			user.Age)
	}
}
