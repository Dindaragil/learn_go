package controllers

import (
	"learn_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all books
func FindBooks(c *gin.Context) {
	database := models.ConnectDatabase()
	defer models.Closedatabase(database)

	var books []models.Book
	database.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

//Create new book
func CreateBook(c *gin.Context) {
	database := models.ConnectDatabase()
	defer models.Closedatabase(database)

	//Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author, Quantity: input.Quantity, IsAvailable: input.IsAvailable}
	database.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

//Get Detail Book
func FindBook(c *gin.Context) {
	database := models.ConnectDatabase()
	defer models.Closedatabase(database)

	// Get model if exist
	var book models.Book

	err := database.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	database := models.ConnectDatabase()
	defer models.Closedatabase(database)
	//Get model if exist
	var book models.Book
	err := database.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	database.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateBook(c *gin.Context) {
	database := models.ConnectDatabase()
	defer models.Closedatabase(database)

	var book models.Book
	if error := database.Where("id = ?", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateBookInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	database.Model(&book).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
