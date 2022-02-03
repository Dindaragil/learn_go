package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

func (book *Book) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("id", uuid.NewV4().String())
	return nil
}
