package main

import (
	"KimLikDogrulaması/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	router.Use(gin.Logger())

	routes.GirisRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	router.Run(":8080")

}
