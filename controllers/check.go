package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck Handler
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": "Testing api", "alive": true})
}
