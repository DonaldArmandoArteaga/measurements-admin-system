package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHealthCheckController(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})
}
