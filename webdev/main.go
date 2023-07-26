package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Book represents the book model
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	router := gin.Default()

	router.GET("/ping", Pong)
	router.GET("/books", GetBooks)
	router.POST("/books", CreateBook)
	router.PUT("/books/:id", UpdateBook)
	router.GET("/books/:id", GetBook)

	fmt.Println("Starting the server on port 8080...")
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Printf("Cannot start the server: %v", err)
		return
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var newBook Book
	err := c.BindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot create book")
		return
	}

	newBook.ID = strconv.Itoa(len(books) + 1)
	books = append(books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var updateBook Book

	err := c.BindJSON(&updateBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot update the book")
		return
	}

	for index, book := range books {
		if book.ID == id {
			books[index] = updateBook
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, "cannot find the book")

}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, "cannot find the book")
}
