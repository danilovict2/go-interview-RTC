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
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	StartTime    time.Time  `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	Status       Status     `json:"status"`
	StreamCallID string     `json:"stream_call_id"`
	Attendees    []User     `gorm:"many2many:interview_attendees"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
