package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.Default()
}

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// You can add more advanced logging here
		c.Next()
	}
}
