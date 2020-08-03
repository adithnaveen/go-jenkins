package main

import (
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy, testing jenkins",
	})
}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", HomePage)

	router.Run("0.0.0.0:5757") //ipaddress:8080/ping
}
