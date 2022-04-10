package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/api-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db not connected")
	}

	// Crete table using migration
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	// Create book
	bookRequest := book.BookRequest{
		Title: "Malam minggu miko",
		Price: 23,
	}

	bookService.CreateBook(bookRequest)

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

	router.Run(":8888")
	// router.Run() // default port 8080

	// workflow : main -> handler -> service -> repository -> db -> mysql
}
