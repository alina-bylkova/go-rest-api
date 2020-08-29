package api

import (
	"go-rest-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Example without interface
// func NumbersGet(db *slice_db.SliceDb) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		result := db.GetAll()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

// Using interface
func NumbersGet(db db.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := db.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
