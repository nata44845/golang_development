package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // sqlite driver
)

func main() {
	// Открытие соединения с базой
	db, err := sql.Open("sqlite", "file:mydb.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
		)`,
	)
	_, err = db.Exec("INSERT INTO users (name, age) VALUES(?, ?)", "John", 21)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name, age) VALUES(?, ?)", "Nik", 21)
	if err != nil {
		log.Fatal(err)
	}
	// Execute a query
	rows, err := db.Query("SELECT id, name, age FROM users WHERE age = ?", 21)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// Пробегаемся по всем строкам, которые получаем из базы
	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s %d\n", id, name, age)
	}

	// Проверяем есть ли ошибка в возвращаемом наборе данных
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
