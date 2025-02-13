package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/GetStream/getstream-go"
	"github.com/danilovict2/go-interview-RTC/models"
	uuidUtils "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (cfg *APIConfig) InterviewStore(c echo.Context) error {
	uuid := uuidFromJWT(c)
	user := models.User{}
	err := cfg.DB.First(&user, "uuid = ?", uuid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "User not found",
		})
	} else if err != nil {
		return HandleGracefully(err, c)
	}

	if user.Role != models.ROLE_INTERVIEWER {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Candidates can not create interviews",
		})
	}

	startTime, err := time.Parse(http.TimeFormat, c.FormValue("startTime"))
	if err != nil {
		return HandleGracefully(err, c)
	}

	status := models.STATUS_LIVE
	if startTime.After(time.Now()) {
		status = models.STATUS_UPCOMING
	}

	callID := uuidUtils.NewString()
	call := cfg.StreamClient.Video().Call("default", callID)
	callRequest := getstream.GetOrCreateCallRequest{
		Data: &getstream.CallRequest{
			CreatedByID: getstream.PtrTo(user.UUID.String()),
		},
	}

	resp, err := call.GetOrCreate(c.Request().Context(), &callRequest)
	if err != nil {
		return HandleGracefully(err, c)
	}

	fmt.Println(resp.Data.Call)

	interview := models.Interview{
		Title:        c.FormValue("title"),
		Description:  c.FormValue("description"),
		StartTime:    startTime,
		Status:       status,
		StreamCallID: resp.Data.Call.ID,
		Attendees:    []models.User{user},
	}

	/*if err := cfg.DB.Create(&interview).Error; err != nil {
		return HandleGracefully(err, c)
	}*/

	return c.JSON(http.StatusOK, interview)
}
