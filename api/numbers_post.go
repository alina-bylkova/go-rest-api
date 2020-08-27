package api

import (
	"go-rest-api/db/slice_db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type numbersPostRequest struct {
	Num int `json:"num"`
}

func NumbersPost(db *slice_db.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := numbersPostRequest{}
		c.Bind(&requestBody)
		result := db.Add(requestBody.Num)
		if result == "Number already exist in the database" {
			c.JSON(http.StatusUnprocessableEntity, result)
		} else {
			c.Status(http.StatusCreated)
			// c.JSON(http.StatusCreated, result)
		}
	}
}
