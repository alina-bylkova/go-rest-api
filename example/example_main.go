package example

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Standard way
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Way using helper function
	//r.GET("/hello", HelloGet)

	// Another way using helper function that return a function
	r.GET("/hello", HelloGet())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
