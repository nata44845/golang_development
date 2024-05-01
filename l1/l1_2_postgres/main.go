package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	db, err := sql.Open("postgres",
		"postgres://postgres:123@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Создание таблицы
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	age INTEGER NOT NULL
	)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// Добавление пользователя
	_, err = db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", "John Doe", 25)
	if err != nil {
		log.Fatal(err)
	}
	// lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Last Insert ID: %d\n", lastInsertId)
	// Получение списка пользователей
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
