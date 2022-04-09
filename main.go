package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)

	router.GET("/book/:id", bookHandler)
	// url value from parameter
	// ex : localhost:8888/book/10

	router.GET("/query", queryHandler)
	// url value from query
	// ex : localhost:8888/query?title=bumi&price=100

	router.Run(":8888") // default port 8080

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tittle": "Hello world",
		"sub":    "Learning Golang",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Juang Sabit",
		"bio":  "do your best",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func queryHandler(c *gin.Context) {
	priceString := c.Query("price")
	title := c.Query("title")
	priceInt, err := strconv.Atoi(priceString)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if priceInt <= 10 {
		status = "cheap"
	} else {
		status = "expensive"
	}
	c.JSON(http.StatusOK, gin.H{
		"title":  title,
		"price":  priceString,
		"status": status,
	})
}
