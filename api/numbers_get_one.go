package api

import (
	"go-rest-api/db/slice_db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NumbersGetOne(db *slice_db.SliceDb) gin.HandlerFunc {
	return func(c *gin.Context) {
		numString := c.Param("int")
		num, _ := strconv.Atoi(numString)
		result, err := db.GetOne(num)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
