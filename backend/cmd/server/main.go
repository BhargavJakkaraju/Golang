package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //creates router with Logger and recovery middleware
	
	r.GET("/Ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	r.Run(":8080")
}