package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`    // space sensitive
	Rating      int    `json:"rating"`   // space sensitive
	Discount    int    `json:"discount"` // space sensitive
}
