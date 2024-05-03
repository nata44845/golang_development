package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	users := []User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "John Doe", Email: "john.doe@example.com"},
		{ID: "3", Name: "Jane Doe", Email: "jane.doe@example.com"},
	}
	// Сохранение пользователей в Redis с использованием HSET
	for _, user := range users {
		key := "user:" + user.ID
		err := rdb.HSet(ctx, key, map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		}).Err()
		if err != nil {
			log.Fatalf("Failed to set user: %v", err)
		}
	}
	fmt.Println("Users saved to Redis")
	// Извлечение всех пользователей и фильтрация по имени "John Doe"
	var filteredUsers []User
	for _, user := range users {
		key := "user:" + user.ID
		result, err := rdb.HGetAll(ctx, key).Result()
		if err != nil {
			log.Fatalf("Failed to get user: %v", err)
		}
		if result["name"] == "John Doe" {
			filteredUsers = append(filteredUsers, User{
				ID:    user.ID,
				Name:  result["name"],
				Email: result["email"],
			})
		}
	}
	fmt.Println("Filtered users:")
	for _, user := range filteredUsers {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", user.ID, user.Name,
			user.Email)
	}
}
