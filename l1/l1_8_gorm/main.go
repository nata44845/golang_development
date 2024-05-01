package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	// Миграция схемы
	db.AutoMigrate(&User{})
	// Добавление нескольких пользователей
	usernames := []string{"John Doe", "Jane Doe"}
	ages := []int{25, 21}
	for i, username := range usernames {
		db.Create(&User{Name: username, Age: ages[i]})
	}
	// Получение списка пользователей с возрастом 25 лет
	var users []User
	db.Where("age = ?", 25).Find(&users)
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name,
			user.Age)
	}
}
