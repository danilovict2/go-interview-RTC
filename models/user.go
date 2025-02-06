package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID      uuid.UUID `gorm:"primaryKey" json:"uuid"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password  []byte    `validate:"required,min=8" json:"-"`
	Role      string    `json:"role" gorm:"default:candidate"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
