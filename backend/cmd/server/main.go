package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //creates router with Logger and recovery middleware
	
	r.GET("/Ping", func(c *gin.Context) { //registers route handler, anonymous function with a pointer to a context obkect
		c.JSON(200, gin.H{ //converts to json format
			"message": "Pong",
		})
	})

    r.POST("/greet", func(c *gin.Context) {
        var data map[string]string
        c.BindJSON(&data)

        name := data["name"]
        c.JSON(200, gin.H{
            "message": "Hello " + name,
        })

    })

    r.Run(":8080")
}