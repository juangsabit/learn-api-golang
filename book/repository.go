package book

import "gorm.io/gorm"

type Repository interface {
	findAll() ([]Book, error)
	findByID(ID int) (Book, error)
	CreateBook(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book
	err := r.db.Take(&book, ID).Error
	return book, err
}

func (r *repository) CreateBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}
