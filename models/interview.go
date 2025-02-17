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

type Decision string

const (
	DECISION_PASS      Decision = "pass"
	DECISION_FAIL      Decision = "fail"
	DECISION_UNDECIDED Decision = "undecided"
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
	Decision     Decision   `json:"decision"  gorm:"default:undecided"`
}
