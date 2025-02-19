package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danilovict2/go-interview-RTC/internal/repository"
	"github.com/danilovict2/go-interview-RTC/internal/services"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (cfg *APIConfig) InterviewStore(c echo.Context) error {
	r := repository.NewUserRepository(cfg.DB)
	user, err := r.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return handleGormError(err, "User", c)
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

func (cfg *APIConfig) InterviewPatch(c echo.Context) error {
	ur := repository.NewUserRepository(cfg.DB)
	user, err := ur.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return handleGormError(err, "User", c)
	}

	ir := repository.NewInterviewRepository(cfg.DB)
	interview, err := ir.FindOneByStreamCallID(c.Param("stream-call-id"))
	if err != nil {
		return handleGormError(err, "Interview", c)
	}

	if !canAccess(user, interview, ir) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "You can not modify this resource",
		})
	}

	if title := c.FormValue("title"); title != "" {
		interview.Title = title
	}

	if description := c.FormValue("description"); description != "" {
		interview.Description = description
	}

	if startTimeStr := c.FormValue("startTime"); startTimeStr != "" {
		startTime, err := time.Parse(http.TimeFormat, startTimeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid Start Time",
			})
		}
		interview.StartTime = startTime
	}

	if endTimeStr := c.FormValue("endTime"); endTimeStr != "" {
		endTime, err := time.Parse(http.TimeFormat, endTimeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid End Time",
			})
		}
		interview.EndTime = &endTime
	}

	if status := models.Status(c.FormValue("status")); status != "" {
		switch status {
		case models.STATUS_UPCOMING, models.STATUS_LIVE, models.STATUS_COMPLETED:
			interview.Status = status
		default:
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid status value",
			})
		}
	}

	if decision := models.Decision(c.FormValue("decision")); decision != "" {
		switch decision {
		case models.DECISION_PASS, models.DECISION_FAIL, models.DECISION_UNDECIDED:
			interview.Decision = decision
		default:
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid decision value",
			})
		}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(interview); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := cfg.DB.Save(&interview).Error; err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, interview)
}

func canAccess(user models.User, interview models.Interview, ir *repository.InterviewRepository) bool {
	isAttendee, err := ir.IsAttendee(interview, user)
	if err != nil {
		return false
	}

	return user.Role == models.ROLE_INTERVIEWER && isAttendee
}

func (cfg *APIConfig) InterviewsGet(c echo.Context) error {
	r := repository.NewUserRepository(cfg.DB)
	user, err := r.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return handleGormError(err, "User", c)
	}

	var interviews []models.Interview
	if err := cfg.DB.Preload("Attendees").Model(&user).Association("Interviews").Find(&interviews); err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"interviews": interviews,
	})
}

func (cfg *APIConfig) InterviewGetComments(c echo.Context) error {
	ur := repository.NewUserRepository(cfg.DB)
	user, err := ur.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return handleGormError(err, "User", c)
	}

	ir := repository.NewInterviewRepository(cfg.DB)
	interview, err := ir.FindOneByStreamCallID(c.Param("stream-call-id"))
	if err != nil {
		return handleGormError(err, "Interview", c)
	}

	if !canAccess(user, interview, ir) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "You can not access this resource",
		})
	}

	comments := make([]models.Comment, 0)
	if err := cfg.DB.Where("interview_id = ?", interview.ID).Preload("CreatedBy").Find(&comments).Error; err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"comments": comments,
	})
}
