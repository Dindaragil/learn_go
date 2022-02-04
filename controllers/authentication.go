package controllers

import (
	"learn_go/helper"
	"learn_go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	database := models.ConnectDatabase()
	defer models.Closedatabase(database)

	// validate input
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//check if the email is already registered
	if err := database.Where("email = ?", input.Email).First(&input).Error; err == nil {
		log.Fatalln("This is not meant to be")
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email is already registered!"})
		return
	}

	var err error
	input.Password, err = helper.GeneratehashPassword(input.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user
	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password, Role: input.Role}
	database.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Sign up successfully!"})
}
