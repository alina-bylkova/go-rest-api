package api

import (
	"go-rest-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Without interface
// func NumbersGetOne(db *slice_db.SliceDb) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		numString := c.Param("int")
// 		num, _ := strconv.Atoi(numString)
// 		result, err := db.GetOne(num)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, err)
// 		} else {
// 			c.JSON(http.StatusOK, result)
// 		}
// 	}
// }

// Using interface
func NumbersGetOne(db db.DataLayer) gin.HandlerFunc {
	return func(c *gin.Context) {
		numString := c.Param("int")
		num, _ := strconv.Atoi(numString)
		result, err := db.GetOne(num)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
