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

	// Create
	// book := book.Book{}
	// book.Title = "Atomic habits"
	// book.Description = "Buku self development"
	// book.Discount = 0
	// book.Price = 10
	// book.Rating = 3

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=======================")
	// 	fmt.Println("Error creating new book")
	// 	fmt.Println("=======================")
	// }

	// Read
	// var books []book.Book
	// // check documentation https://gorm.io/docs/query.html
	// err = db.Debug().Where("rating > ?", 1).Find(&books).Error // Debug() = get query sql;
	// if err != nil {
	// 	fmt.Println("=======================")
	// 	fmt.Println(" Error get data ")
	// 	fmt.Println("=======================")
	// }

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// 	fmt.Println("Description :", book.Description)
	// 	fmt.Println("Book object %v", book)
	// }

	// Update
	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).Take(&book).Error // Debug() = get query sql;
	// if err != nil {
	// 	fmt.Println("================")
	// 	fmt.Println(" Error get data ")
	// 	fmt.Println("================")
	// }

	// book.Title = "Laskar Pelangi"
	// err = db.Debug().Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=================")
	// 	fmt.Println(" Error save data ")
	// 	fmt.Println("=================")
	// }

	// Delete
	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).Take(&book).Error // Debug() = get query sql;
	// if err != nil {
	// 	fmt.Println("================")
	// 	fmt.Println(" Error get data ")
	// 	fmt.Println("================")
	// }

	// err = db.Debug().Delete(&book).Error // db.Delete() = update field deleted_at
	// if err != nil {
	// 	fmt.Println("=================")
	// 	fmt.Println(" Error save data ")
	// 	fmt.Println("=================")
	// }

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

}
