package api

import (
	"go-rest-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Without interface
// func NumbersDelete(db *slice_db.SliceDb) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		numString := c.Param("int")
// 		num, _ := strconv.Atoi(numString)
// 		result := db.Delete(num)
// 		if !result {
// 			c.JSON(http.StatusNotFound, "Number is not found")
// 		} else {
// 			c.Status(http.StatusNoContent)
// 		}
// 	}
// }

// Using interface
func NumbersDelete(db db.DataLayer) gin.HandlerFunc {
	return func(c *gin.Context) {
		numString := c.Param("int")
		num, _ := strconv.Atoi(numString)
		result := db.Delete(num)
		if !result {
			c.JSON(http.StatusNotFound, "Number is not found")
			return
		}
		c.Status(http.StatusNoContent)
	}
}
