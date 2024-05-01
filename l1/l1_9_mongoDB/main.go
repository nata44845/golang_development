package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// Установка контекста выполнения
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Подключение к MongoDB
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	// Получение ссылки на коллекцию 'users'
	collection := client.Database("testdb").Collection("users")
	// Добавление нового пользователя
	newUser := User{Name: "John Doe", Age: 30}
	insertResult, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user with ID:", insertResult.InsertedID)
	// Получение и вывод всех пользователей
	var users []User
	{
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		if err = cursor.All(ctx, &users); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Found users:")
		for _, user := range users {
			fmt.Println(user)
		}
	}
	{
		// Способ итерации по одному элементу вручную
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		for cursor.Next(ctx) {
			var u User
			if err := cursor.Decode(&u); err != nil {
				return
			}
			users = append(users, u)
		}
		fmt.Println("Found users:")
		for _, user := range users {
			fmt.Println(user)
		}
	}
}
