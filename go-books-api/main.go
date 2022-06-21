package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/books", listBooksHandler)
	r.POST("/books", createBookHandler)
	r.PATCH("/books/:id", updateBookHandler)
	r.DELETE("/books/:id", deleteBookHandler)

	r.Run()
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func listBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func updateBookHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	// Find book by id
	var booksMap = map[string]Book{}
	for _, book := range books {
		booksMap[book.ID] = book
	}
	updateBook, exists := booksMap[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Validate book
	var updateBookInput UpdateBookInput
	if err := c.ShouldBindJSON(&updateBookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update book
	updateBook.Title = updateBookInput.Title
	updateBook.Author = updateBookInput.Author
	for index, book := range books {
		if book.ID == id {
			books[index].Author = updateBookInput.Author
			books[index].Title = updateBookInput.Title
			break
		}
	}

	c.JSON(http.StatusOK, updateBook)
}

func deleteBookHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
