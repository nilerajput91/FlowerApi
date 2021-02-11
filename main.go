package main

import (
	"fmt"

	"github.com/flowerapi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Flower Api")
	router := SetupRouter()
	router.Run(":8081")
}

// SetupRouter
func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("/flower", controllers.Create)
		v1.GET("/flower/:id", controllers.GetFlower)
		v1.GET("/flowers", controllers.GetAllFlowers)
		v1.DELETE("/flower", controllers.DeleteFlower)
		v1.GET("/check", controllers.HealthCheck)
	}
	return router
}
