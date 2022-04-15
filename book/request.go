package book

import "encoding/json"

type BookRequest struct {
	Title       interface{} `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required"` // space sensitive
	Description interface{} `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required"`   // space sensitive
	Discount    json.Number `json:"discount" binding:"required"` // space sensitive
}

type UpdateBookRequest struct {
	Title       interface{} `json:"title"`
	Price       json.Number `json:"price"` // space sensitive
	Description interface{} `json:"description"`
	Rating      json.Number `json:"rating"`   // space sensitive
	Discount    json.Number `json:"discount"` // space sensitive
}
