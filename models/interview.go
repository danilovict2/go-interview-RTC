package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	STATUS_LIVE      Status = "live"
	STATUS_UPCOMING  Status = "upcoming"
	STATUS_COMPLETED Status = "completed"
)

type Interview struct {
	gorm.Model
	Title        string     `json:"title" validate:"required"`
	Description  string     `json:"description" validate:"required"`
	StartTime    time.Time  `json:"start_time" validate:"required"`
	EndTime      *time.Time `json:"end_time"`
	Status       Status     `json:"status"`
	StreamCallID string     `json:"stream_call_id"`
	Attendees    []User     `json:"attendees" gorm:"many2many:interview_attendees" validate:"required,min=1"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
