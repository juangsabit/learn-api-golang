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

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tittle": "Hello world",
		"sub":    "Learning Golang",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Juang Sabit",
		"bio":  "do your best",
	})
}

func (h *bookHandler) BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
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

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}
	bookResponse := convertToBookRespons(b)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookRespons(b)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
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

	book, err := h.bookService.CreateBook(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookRespons(book),
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var updateBookRequest book.UpdateBookRequest
	err := c.ShouldBindJSON(&updateBookRequest)
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

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.UpdateBook(id, updateBookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success updated book",
		"data":   convertToBookRespons(book),
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}
	bookResponse := convertToBookRespons(b)
	c.JSON(http.StatusOK, gin.H{
		"status": "success deleted book",
		"data":   bookResponse,
	})
}

func convertToBookRespons(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          int(b.ID),
		Title:       b.Title,
		Description: b.Description,
		Price:       int(b.Price),
		Rating:      int(b.Rating),
		Discount:    int(b.Discount),
	}
}
