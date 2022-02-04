package middleware

import (
	"learn_go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized or Invalid Token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
