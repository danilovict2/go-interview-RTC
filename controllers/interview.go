package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/danilovict2/go-interview-RTC/internal/repository"
	"github.com/danilovict2/go-interview-RTC/internal/services"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (cfg *APIConfig) InterviewStore(c echo.Context) error {
	r := repository.NewUserRepository(cfg.DB)
	user, err := r.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return HandleGracefully(err, c)
	}

	if user.Role != models.ROLE_INTERVIEWER {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Candidates can not create interviews",
		})
	}

	var attendeeUUIDs []string
	if err := json.Unmarshal([]byte(c.FormValue("attendeeUUIDs")), &attendeeUUIDs); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	attendees, err := r.FindByUUID(attendeeUUIDs)
	if err != nil {
		return HandleGracefully(err, c)
	}

	startTime, err := time.Parse(http.TimeFormat, c.FormValue("startTime"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid Start Time",
		})
	}

	status := models.STATUS_LIVE
	switch {
	case startTime.After(time.Now()):
		status = models.STATUS_UPCOMING
	case startTime.Before(time.Now()):
		status = models.STATUS_COMPLETED
	}

	streamCallID, err := services.CreateNewVideoCall(cfg.StreamClient, user, startTime)
	if err != nil {
		return HandleGracefully(err, c)
	}

	interview := models.Interview{
		Title:        c.FormValue("title"),
		Description:  c.FormValue("description"),
		StartTime:    startTime,
		Status:       status,
		StreamCallID: streamCallID,
		Attendees:    attendees,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(interview); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := cfg.DB.Create(&interview).Error; err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, interview)
}

func (cfg *APIConfig) InterviewEnd(c echo.Context) error {
	ur := repository.NewUserRepository(cfg.DB)
	user, err := ur.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return HandleGracefully(err, c)
	}

	ir := repository.NewInterviewRepository(cfg.DB)
	interview, err := ir.FindOneByStreamCallID(c.Param("stream-call-id"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Interview not found",
		})
	} else if err != nil {
		return HandleGracefully(err, c)
	}

	isAttendee, err := ir.IsAttendee(interview, user)
	if err != nil {
		return HandleGracefully(err, c)
	}

	if user.Role != models.ROLE_INTERVIEWER || !isAttendee {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "You can not modify this resource",
		})
	}

	endTime, err := time.Parse(http.TimeFormat, c.FormValue("endTime"))
	if err != nil {
		return HandleGracefully(err, c)
	}

	interview.Status = models.STATUS_COMPLETED
	interview.EndTime = &endTime

	if err := cfg.DB.Save(&interview).Error; err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, interview)
}

func (cfg *APIConfig) InterviewsGet(c echo.Context) error {
	r := repository.NewUserRepository(cfg.DB)
	user, err := r.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return HandleGracefully(err, c)
	}

	var interviews []models.Interview
	if err := cfg.DB.Model(&user).Association("Interviews").Find(&interviews); err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"interviews": interviews,
	})
}
