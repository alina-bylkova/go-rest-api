package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Helper function
// func HelloGet(c *gin.Context) {
// 	c.JSON(http.StatusOK, map[string]string{
// 		"hello": "Found me",
// 	})
// }

// Helper function that returns a function
func HelloGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
