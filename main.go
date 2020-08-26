package main

import (
	"go-rest-api/api"
	"go-rest-api/db/slice_db"

	"github.com/gin-gonic/gin"
)

func main() {
	api.Hello()

	r := gin.Default()

	r.GET("/hello", api.HelloGet())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	db := slice_db.New()

	db.Add(25)
}
