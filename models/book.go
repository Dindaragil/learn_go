package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Quantity    uint       `json:"quantity"`
	IsAvailable bool       `json:"isAvailable"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
}

type CreateBookInput struct {
	Title       string `json:"title" binding: "required"`
	Author      string `json:"author" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	IsAvailable bool   `json:"isAvailable" binding:"required"`
}

type UpdateBookInput struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Quantity    uint   `json:"quantity"`
	IsAvailable bool   `json:"isAvailable"`
}

func (book *Book) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("id", uuid.NewV4().String())
	return nil
}
