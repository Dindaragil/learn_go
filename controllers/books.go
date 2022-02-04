package controllers

import (
	"learn_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

//Create new book
func CreateBook(c *gin.Context) {
	//Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author, Quantity: input.Quantity, IsAvailable: input.IsAvailable}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

//Get Detail Book
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book

	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	//Get model if exist
	var book models.Book
	err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if error := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&book).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
