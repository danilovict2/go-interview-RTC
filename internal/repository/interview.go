package repository

import (
	"github.com/danilovict2/go-interview-RTC/models"
	"gorm.io/gorm"
)

type InterviewRepository struct {
	DB *gorm.DB
}

func NewInterviewRepository(DB *gorm.DB) *InterviewRepository {
	return &InterviewRepository{
		DB: DB,
	}
}

func (r *InterviewRepository) FindOneByStreamCallID(streamCallID string) (models.Interview, error) {
	interview := models.Interview{}
	if err := r.DB.First(&interview, "stream_call_id = ?", streamCallID).Error; err != nil {
		return models.Interview{}, err
	}

	return interview, nil
}

func (r *InterviewRepository) IsAttendee(interview models.Interview, user models.User) (bool, error) {
	attendees := make([]models.User, 0)
	if err := r.DB.Model(&interview).Association("Attendees").Find(&attendees); err != nil {
		return false, err
	}

	for _, attendee := range attendees {
		if attendee.UUID == user.UUID {
			return true, nil
		}
	}

	return false, nil
}
