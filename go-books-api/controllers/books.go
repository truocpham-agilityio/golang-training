package controllers

import (
	"example/go-books-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /books
func FindBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
func CreateBook(c *gin.Context) {
	// Validate input
	var createBookInput CreateBookInput
	if err := c.ShouldBindJSON((&createBookInput)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: createBookInput.Title, Author: createBookInput.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// GET /books/:id
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	// Validate input
	var updateBookInput UpdateBookInput
	if err := c.ShouldBindJSON(&updateBookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update model
	models.DB.Model(&book).Updates(updateBookInput)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	// Delete model
	models.DB.Delete(&book)

	c.Status(http.StatusNoContent)
}