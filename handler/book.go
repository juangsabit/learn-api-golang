package handler

import (
	"fmt"
	"log"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tittle": "Hello world",
		"sub":    "Learning Golang",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Juang Sabit",
		"bio":  "do your best",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func QueryHandler(c *gin.Context) {
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

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookRequest
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
