package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	var (
		method string
		id     int
		title  string
		author string
	)
	flag.StringVar(&method, "method", "get", "HTTP method (get, post, put,delete)")
	flag.IntVar(&id, "id", 0, "Book ID")
	flag.StringVar(&title, "title", "", "Book title")
	flag.StringVar(&author, "author", "", "Book author")
	flag.Parse()
	var url string
	if method == "get" && id == 0 {
		url = "http://localhost:8080/books"
	} else if method == "get" && id != 0 {
		url = fmt.Sprintf("http://localhost:8080/books/%d", id)
	} else if method == "post" || method == "put" {
		url = "http://localhost:8080/books"
	} else if method == "delete" {
		url = fmt.Sprintf("http://localhost:8080/books/%d", id)
	} else {
		fmt.Println("Неверный метод или параметры")
		return
	}
	var req *http.Request
	var err error
	if method == "post" || method == "put" {
		book := Book{ID: id, Title: title, Author: author}
		jsonData, _ := json.Marshal(book)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}
	fmt.Println(string(body))
}
