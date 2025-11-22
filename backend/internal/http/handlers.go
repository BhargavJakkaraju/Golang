package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//healthCheck route

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status" : "Ok",
	})
}

func GetSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Active",
	})
}
