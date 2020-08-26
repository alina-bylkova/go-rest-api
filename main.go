package main

import (
	"go-rest-api/api"

	"github.com/gin-gonic/gin"
)

func main() {
	api.Hello()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// One way
	// r.GET("/hello", api.HelloGet)

	// Another way
	r.GET("/hello", api.HelloGet())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
