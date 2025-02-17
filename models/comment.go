package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content       string    `json:"content" gorm:"type:text"`
	Rating        int       `json:"rating"`
	CreatedByUUID uuid.UUID `json:"-"`
	CreatedBy     User      `json:"created_by" gorm:"foreignKey:CreatedByUUID"`
	InterviewID   uint      `json:"-"`
	Interview     Interview `json:"interview"`
}
