package main

import (
	"concurrent-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/concurrent", handlers.ConcurrentFetch)

	r.Run(":8080") // http://localhost:8080/concurrent
}
