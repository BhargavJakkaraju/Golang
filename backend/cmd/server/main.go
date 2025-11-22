package main

import (
	"golang/backend/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //creates router with Logger and recovery middleware
    http.RegisterRoute(r)
    r.Use(http.LoggerMiddleaware())
    r.Run(":8080")

}