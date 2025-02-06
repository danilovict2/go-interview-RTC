package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID       uuid.UUID    `json:"uuid" gorm:"primaryKey"`
	FirstName  string       `json:"first_name" validate:"required"`
	LastName   string       `json:"last_name" validate:"required"`
	Email      string       `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	Password   []byte       `json:"-" validate:"required,min=8"`
	Role       string       `json:"role" gorm:"default:candidate"`
	Interviews *[]Interview `json:"interviews" gorm:"many2many:interview_attendees"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
