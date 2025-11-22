package http

import (
	"github.com/gin-gonic/gin"
)

func LoggerMiddleaware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Before Request", c.Request.URL.Path)
		c.Next()
		println("After Request")
	}
}