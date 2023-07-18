package main

import (
	"fmt"
	"net/http"
	"webdev/book"

	"github.com/gorilla/mux"
)

func main() {
	book.InitializeBook()

	router := mux.NewRouter()
	router.HandleFunc("/books", book.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", book.GetBook).Methods(http.MethodGet)

	fmt.Println("Starting the server at port: ", 8080)

	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		fmt.Println("Cannot start the server: ", err)
	}
}
