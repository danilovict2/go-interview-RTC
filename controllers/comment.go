package controllers

import (
	"net/http"
	"strconv"

	"github.com/danilovict2/go-interview-RTC/internal/repository"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/labstack/echo/v4"
)

func (cfg *APIConfig) CommentStore(c echo.Context) error {
	ur := repository.NewUserRepository(cfg.DB)
	user, err := ur.FindOneByUUID(c.Get("uuid").(string))
	if err != nil {
		return handleGormError(err, "User", c)
	}

	if user.Role != models.ROLE_INTERVIEWER {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Candidates can not create comments",
		})
	}

	ir := repository.NewInterviewRepository(cfg.DB)
	interview, err := ir.FindOneByStreamCallID(c.FormValue("interviewID"))
	if err != nil {
		return handleGormError(err, "Interview", c)
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid rating",
		})
	}

	comment := models.Comment{
		Content:       c.FormValue("content"),
		Rating:        rating,
		CreatedByUUID: user.UUID,
		CreatedBy:     user,
		InterviewID:   interview.ID,
		Interview:     interview,
	}

	if err := cfg.DB.Create(&comment).Error; err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"comment": comment,
	})
}
