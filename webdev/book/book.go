package book

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var nextID = 1

func InitializeBook() {
	b1 := Book{
		ID:     nextID,
		Title:  "Book 1",
		Author: "Author 1",
	}
	books = append(books, b1)
	nextID++

	b2 := Book{
		ID:     nextID,
		Title:  "Book 2",
		Author: "Author 2",
	}
	books = append(books, b2)
	nextID++
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bx, err := json.Marshal(books)
	if err != nil {
		fmt.Println("cannot marshal books")
		return
	}

	w.Write(bx)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bx, err := json.Marshal(books)
	if err != nil {
		fmt.Println("cannot marshal books")
		return
	}

	w.Write(bx)
}
