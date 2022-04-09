package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/book/:id", handler.BookHandler)
	// url value from parameter
	// ex : localhost:8888/book/10

	v1.GET("/query", handler.QueryHandler)
	// url value from query
	// ex : localhost:8888/query?title=bumi&price=100

	v1.POST("/book", handler.PostBookHandler)

	router.Run(":8888") // default port 8080

}
