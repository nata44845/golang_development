package client

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	var books []server.Book
	err = client.Call("BookService.GetBooks", struct{}{}, &books)
	if err != nil {
		log.Fatalf("Failed to call GetBooks: %v", err)
	}
	// Обработка полученных книг
	for _, book := range books {
		log.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title,
			book.Author)
	}
	// Вызов других методов аналогичным образом
	var book Book
	err = client.Call("BookService.GetBook", 1, &book)
	if err != nil {
		log.Fatalf("Failed to call GetBook: %v", err)
	}
	log.Printf("Book: %+v\n", book)
	// Добавление новой книги
	newBook := Book{ID: 3, Title: "Book 3", Author: "Author 3"}
	err = client.Call("BookService.CreateBook", &newBook, &newBook)
	if err != nil {
		log.Fatalf("Failed to call CreateBook: %v", err)
	}
	log.Printf("New book created: %+v\n", newBook)
	// Обновление книги
	updatedBook := Book{ID: 1, Title: "Updated Book 1", Author: "Updated Author 1"}
	err = client.Call("BookService.UpdateBook", &updatedBook, &updatedBook)
	if err != nil {
		log.Fatalf("Failed to call UpdateBook: %v", err)
	}
	log.Printf("Book updated: %+v\n", updatedBook)
	// Удаление книги
	err = client.Call("BookService.DeleteBook", 1, nil)
	if err != nil {
		log.Fatalf("Failed to call DeleteBook: %v", err)
	}
	log.Println("Book deleted")
}
