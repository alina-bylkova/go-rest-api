package api

import (
	"go-rest-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type numbersPutRequest struct {
	NewNum int `json:"newNum"`
}

// Using interface and Update function in db package
// func NumbersUpdate(db db.DataLayer) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestBody := numbersPutRequest{}
// 		c.Bind(&requestBody)

// 		numString := c.Param("int")
// 		num, _ := strconv.Atoi(numString)

// 		result := db.Update(num, requestBody.NewNum)

// 		if result != nil {
// 			c.JSON(http.StatusUnprocessableEntity, result)
// 			return
// 		}
// 		c.Status(http.StatusNoContent)
// 	}
// }

// Updated version
func NumbersUpdate(db db.DataLayer) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := numbersPutRequest{}
		c.Bind(&requestBody)

		numString := c.Param("int")
		num, _ := strconv.Atoi(numString)

		if _, err := db.GetOne(requestBody.NewNum); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Already exists")
			return
		}

		if res := db.Delete(num); !res {
			c.JSON(http.StatusUnprocessableEntity, "Searched element not found")
			return
		}

		db.Add(requestBody.NewNum)
		c.Status(http.StatusNoContent)
	}
}
