package api

import (
	"go-rest-api/db/slice_db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NumbersDelete(db *slice_db.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		numString := c.Param("int")
		num, _ := strconv.Atoi(numString)
		result := db.Delete(num)
		if result == "Number is not found" {
			c.JSON(http.StatusNotFound, result)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}
