package book

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(225)"`
	Description string `gorm:"type:varchar(225)"`
	Price       int    `gorm:"type:int(11)"`
	Rating      int    `gorm:"type:int(11)"`
	Discount    int    `gorm:"type:int(11)"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
