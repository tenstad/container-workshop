package main

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		message := os.Getenv("MESSAGE")
		if message == "" {
			message = "Hello, World!"
		}

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
