package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/book/:id", bookHandler)
	// url value from parameter
	// ex : localhost:8888/book/10

	v1.GET("/query", queryHandler)
	// url value from query
	// ex : localhost:8888/query?title=bumi&price=100

	v1.POST("/book", postBookHandler)

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
	priceInt, err := strconv.Atoi(priceString) // convert str to int
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

type BookInput struct {
	Title interface{} `json:"title" binding:"required,lowercase"`
	Price interface{} `json:"price" binding:"required,number"` // space sensitive
}

func postBookHandler(c *gin.Context) {
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
