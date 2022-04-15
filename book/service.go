package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	CreateBook(bookRequest BookRequest) (Book, error)
	UpdateBook(ID int, updateBookRequest UpdateBookRequest) (Book, error)
	DeleteBook(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) CreateBook(bookRequest BookRequest) (Book, error) {
	title, _ := bookRequest.Title.(string)
	price, _ := bookRequest.Price.Int64()
	description, _ := bookRequest.Description.(string)
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       title,
		Price:       int(price),
		Description: description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.CreateBook(book)
	return newBook, err
}

func (s *service) UpdateBook(ID int, updateBookRequest UpdateBookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)

	title, _ := updateBookRequest.Title.(string)
	price, _ := updateBookRequest.Price.Int64()
	description, _ := updateBookRequest.Description.(string)
	rating, _ := updateBookRequest.Rating.Int64()
	discount, _ := updateBookRequest.Discount.Int64()

	book.Title = title
	book.Price = int(price)
	book.Description = description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.UpdateBook(book)
	return newBook, err
}

func (s *service) DeleteBook(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.DeleteBook(book)
	return newBook, err
}
