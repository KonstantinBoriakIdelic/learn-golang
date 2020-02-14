package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"response": "Learn Golang application is running"})
	})

	router.Run(":" + port)
}
