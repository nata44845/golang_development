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

// Определение структуры User, соответствующей документу в MongoDB
type User struct {
	Name    string   `bson:"name"`    // Имя пользователя
	Age     int      `bson:"age"`     // Возраст пользователя
	Hobbies []string `bson:"hobbies"` // Список хобби пользователя
}

func main() {
	// Создание контекста с таймаутом для управления подключением к MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Подключение к серверу MongoDB
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx) // Отложенное отключение от сервера MongoDB
	// Получение ссылки на коллекцию 'users'
	collection := client.Database("testdb").Collection("users")
	// Добавление нескольких пользователей в коллекцию
	users := []interface{}{
		User{"Alice", 30, []string{"reading", "traveling"}},
		User{"Bob", 25, []string{"coding", "gaming"}},
		User{"Charlie", 35, []string{"cooking", "reading"}},
	}
	_, err = collection.InsertMany(ctx, users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users inserted")

	// Агрегационный конвейер для получения уникальных хобби и среднего возраста
	pipeline := mongo.Pipeline{
		{{"$unwind", "$hobbies"}}, // Разделение документов по хобби
		{
			{
				"$group", bson.D{{"_id", "$hobbies"}, {"averageAge",
					bson.D{{"$avg", "$age"}}}},
			},
		}, // Группировка по хобби с вычислением среднего возраста
		{{"$sort", bson.D{{"averageAge", 1}}}}, // Сортировка хобби по среднему возрасту
	}
	// Выполнение агрегационного запроса
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M // Слайс для хранения результатов агрегации
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	// Вывод результатов агрегации
	fmt.Println("Hobbies and their average age of interest:")
	for _, result := range results {
		fmt.Printf("Hobby: %v, Average Age: %v\n", result["_id"],
			result["averageAge"])
	}
}
