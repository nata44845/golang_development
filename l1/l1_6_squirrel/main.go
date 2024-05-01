package main

import (
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
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
	// Добавление нескольких пользователей
	usernames := []string{"John Doe", "Jane Doe"}
	ages := []int{25, 21}
	for i, username := range usernames {
		db.MustExec("INSERT INTO users (name, age) VALUES ($1, $2)",
			username, ages[i])
	}
	// Получение списка пользователей с возрастом 25 лет с использованием squirrel
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select("*").From("users").Where(sq.Eq{"age": 21})

	rows, err := query.RunWith(db).Query()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name,
			user.Age)
	}
}
