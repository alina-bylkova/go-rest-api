package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// One way
// func HelloGet(c *gin.Context) {
// 	c.JSON(http.StatusOK, map[string]string{
// 		"hello": "Found me",
// 	})
// }

// Another way
func HelloGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
