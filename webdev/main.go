package main

import (
	"fmt"
	"webdev2/handler"

	"github.com/gin-gonic/gin"
)

// Book represents the book model

func main() {
	router := gin.Default()

	router.GET("/ping", handler.Pong)
	router.GET("/books", handler.GetBooks)
	router.POST("/books", handler.CreateBook)
	router.PUT("/books/:id", handler.UpdateBook)
	router.GET("/books/:id", handler.GetBook)
	router.DELETE("/books/:id", handler.DeleteBook)

	fmt.Println("Starting the server on port 8080...")
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Printf("Cannot start the server: %v", err)
		return
	}
}
