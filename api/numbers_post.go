package api

import (
	"go-rest-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type numbersPostRequest struct {
	Num int `json:"num"`
}

// Example without interface
// func NumbersPost(db *slice_db.SliceDb) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestBody := numbersPostRequest{}
// 		c.Bind(&requestBody)
// 		result := db.Add(requestBody.Num)
// 		if !result {
// 			c.JSON(http.StatusUnprocessableEntity, "Number already exist in the database")
// 		} else {
// 			c.Status(http.StatusCreated)
// 			// c.JSON(http.StatusCreated, result)
// 		}
// 	}
// }

// Using interface
func NumbersPost(db db.DataLayer) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := numbersPostRequest{}
		c.Bind(&requestBody)
		result := db.Add(requestBody.Num)
		if !result {
			c.JSON(http.StatusUnprocessableEntity, "Number already exist in the database")
			return
		}
		c.Status(http.StatusCreated)
		// c.JSON(http.StatusCreated, result)

	}
}
