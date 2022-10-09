package main

import (
	"ginJwt/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := gin.New()
	r.Use(gin.Logger())

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access granted for api-1",
		})

	})
	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Access granted for api-2",
		})

	})
	r.Run(":" + port)
}
