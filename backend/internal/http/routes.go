package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	router.GET("/health", HealthCheck)
	router.GET("/session", GetSession)

}