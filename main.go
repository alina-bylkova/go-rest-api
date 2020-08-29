package main

import (
	"go-rest-api/api"
	"go-rest-api/db/redis_db"

	"github.com/gin-gonic/gin"
)

func main() {
	// var db *redis_db.RedisDb = redis_db.New()
	// var db *slice_db.SliceDb = slice_db.New()
	// var db db.DataLayer = redis_db.New()
	// var db db.DataLayer = slice_db.New()

	db := redis_db.New()

	r := gin.Default()

	r.GET("/numbers", api.NumbersGet(db))
	r.GET("/numbers/:int", api.NumbersGetOne(db))
	r.POST("/numbers", api.NumbersPost(db))
	r.DELETE("/numbers/:int", api.NumbersDelete(db))
	r.PUT("/numbers/:int", api.NumbersUpdate(db))

	r.Run()
}
