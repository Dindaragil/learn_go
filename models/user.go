package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	Name      string     `json:"name"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type RegisterInput struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("id", uuid.NewV4().String())
	return nil
}
