package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/to4to/Go-JWT-Auth/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access Granted for Api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access Granted for Api-2"})

	})


	router.Run(":"+port)
}
//21.22