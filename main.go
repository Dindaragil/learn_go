package main

import (
	"learn_go/controllers"
	"learn_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
