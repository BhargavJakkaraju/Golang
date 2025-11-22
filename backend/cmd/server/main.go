package main

import (
	"golang/backend/internal/http"
    "golang/backend/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database first
	db.Connect()
	
	// Setup router
	r := gin.Default() //creates router with Logger and recovery middleware
	r.Use(http.LoggerMiddleaware())
	http.RegisterRoute(r)
	r.Run(":8080")
}