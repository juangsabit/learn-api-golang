package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	CreateBook(bookRequest BookRequest) (Book, error)
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
	return s.repository.FindByID(1)
}

func (s *service) CreateBook(bookRequest BookRequest) (Book, error) {
	title, _ := bookRequest.Title.(string)
	price, _ := bookRequest.Price.(int)
	book := Book{
		Title: title,
		Price: price,
	}

	newBook, err := s.repository.CreateBook(book)
	return newBook, err
}
