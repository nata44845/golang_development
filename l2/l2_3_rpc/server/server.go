package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookService struct {
	books []Book
}

func (s *BookService) GetBooks(args *struct{}, reply *[]Book) error {
	*reply = s.books
	return nil
}

func (s *BookService) GetBook(args *int, reply *Book) error {
	for _, book := range s.books {
		if book.ID == *args {
			*reply = book
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

func (s *BookService) CreateBook(args *Book, reply *Book) error {
	s.books = append(s.books, *args)
	*reply = *args
	return nil
}

func (s *BookService) UpdateBook(args *Book, reply *Book) error {
	for i, book := range s.books {
		if book.ID == args.ID {
			s.books[i] = *args
			*reply = *args
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

func (s *BookService) DeleteBook(args *int, reply *struct{}) error {
	for i, book := range s.books {
		if book.ID == *args {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}
func main() {
	bookService := &BookService{
		books: []Book{
			{ID: 1, Title: "Book 1", Author: "Author 1"},
			{ID: 2, Title: "Book 2", Author: "Author 2"},
		},
	}
	rpc.Register(bookService)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", listener.Addr().String())
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
