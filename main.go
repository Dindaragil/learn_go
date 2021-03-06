package main

import (
	"learn_go/controllers"
	"learn_go/middleware"
	"learn_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitialMigration()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// protected := r.Group("/private")
	r.Use(middleware.JwtAuthMiddleware())

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.Run()
}
