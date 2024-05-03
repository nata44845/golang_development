package main

import (
	"context"
	"encoding/json"
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

func serializeUser(user User) (string, error) {
	userData, err := json.Marshal(user)
	return string(userData), err
}
func deserializeUser(userData string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(userData), &user)
	return user, err
}
func main() {
	// Инициализация клиента Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// Создание экземпляра пользователя
	user := User{
		ID:    "1",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	// Сериализация пользователя в JSON
	userJson, err := serializeUser(user)
	if err != nil {
		log.Fatalf("Error serializing user: %v", err)
	}
	// Сохранение сериализованного пользователя в Redis
	err = rdb.Set(ctx, "user:"+user.ID, userJson, 0).Err()
	if err != nil {
		log.Fatalf("Error saving user to Redis: %v", err)
	}
	// Извлечение и десериализация пользователя из Redis
	userJson, err = rdb.Get(ctx, "user:"+user.ID).Result()
	if err != nil {
		log.Fatalf("Error retrieving user from Redis: %v", err)
	}
	retrievedUser, err := deserializeUser(userJson)
	if err != nil {
		log.Fatalf("Error deserializing user: %v", err)
	}
	fmt.Printf("Retrieved User: %+v\n", retrievedUser)
}
