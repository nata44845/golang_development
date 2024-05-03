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

type Book struct {
	Title  string `bson:"title"`
	Author string `bson:"author"`
	Year   int    `bson:"year"`
	Pages  int    `bson:"pages"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("library").Collection("books")
	// Добавление книг в коллекцию
	books := []interface{}{
		Book{"The Hobbit", "J.R.R. Tolkien", 1937, 310},
		Book{"1984", "George Orwell", 1949, 328},
		Book{"Brave New World", "Aldous Huxley", 1932, 311},
		Book{"The Great Gatsby", "F. Scott Fitzgerald", 1925, 218},
	}
	_, err = collection.InsertMany(ctx, books)
	if err != nil {
		log.Fatal(err)
	}
	// Фильтрация книг, выпущенных после 1930 года, с сортировкой по году выпуска
	filter := bson.M{"year": bson.M{"$gt": 1930}}
	findOptions := options.Find().SetSort(bson.D{{"year", 1}}) // 1 для сортировки по возрастанию
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	var filteredBooks []Book
	if err = cursor.All(ctx, &filteredBooks); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Filtered books:")
	for _, book := range filteredBooks {
		fmt.Printf("Title: %s, Author: %s, Year: %d, Pages: %d\n", book.Title,
			book.Author, book.Year, book.Pages)
	}
}
