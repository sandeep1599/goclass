package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
			books[index].ID = id
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

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(http.StatusOK, "Book deleted successfully")
			return
		}
	}

	c.JSON(http.StatusNotFound, "Book not found")
}
