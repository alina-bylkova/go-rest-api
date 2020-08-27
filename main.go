package main

import (
	"go-rest-api/api"
	"go-rest-api/db/slice_db"

	"github.com/gin-gonic/gin"
)

func main() {
	db := slice_db.New()
	r := gin.Default()

	r.GET("/numbers", api.NumbersGet(db))
	r.GET("/numbers/:int", api.NumbersGetOne(db))
	r.POST("/numbers", api.NumbersPost(db))
	r.DELETE("/numbers/:int", api.NumbersDelete(db))

	r.Run()
}
