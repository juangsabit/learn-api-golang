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
	bookHandler := handler.NewBookHandler(bookService)

	// Create book
	// bookRequest := book.BookRequest{
	// 	Title:       "Malam minggu miko",
	// 	Price:       25,
	// 	Description: "example description",
	// 	Rating:      4,
	// 	Discount:    10,
	// }

	// bookService.CreateBook(bookRequest)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)

	v1.GET("/getparam/:id", bookHandler.BookHandler)
	// url value from parameter
	// ex : localhost:8888/book/10

	v1.GET("/query", bookHandler.QueryHandler)
	// url value from query
	// ex : localhost:8888/query?title=bumi&price=100

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.POST("/book", bookHandler.PostBookHandler)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)

	router.Run(":8888")
	// router.Run() // default port 8080

	// workflow : main -> handler -> service -> repository -> db -> mysql
}
