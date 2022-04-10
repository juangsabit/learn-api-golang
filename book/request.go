package book

type BookRequest struct {
	Title interface{} `json:"title" binding:"required,lowercase"`
	Price interface{} `json:"price" binding:"required,number"` // space sensitive
}
