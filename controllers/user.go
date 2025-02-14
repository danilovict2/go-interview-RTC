package controllers

import (
	"errors"
	"net/http"

	"github.com/danilovict2/go-interview-RTC/internal/repository"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (cfg *APIConfig) UserStore(c echo.Context) error {
	user := models.User{
		UUID:      uuid.New(),
		FirstName: c.FormValue("first_name"),
		LastName:  c.FormValue("last_name"),
		Email:     c.FormValue("email"),
		// Store the password temporarily for validation purposes
		Password: []byte(c.FormValue("password")),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user.Password, err = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return HandleGracefully(err, c)
	}

	err = cfg.DB.Create(&user).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return c.String(http.StatusBadRequest, "User with this email already exists.")
	} else if err != nil {
		return HandleGracefully(err, c)
	}

	return cfg.Login(c)
}

func (cfg *APIConfig) UserGet(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid == "me" {
		uuid = c.Get("uuid").(string)
	}

	r := repository.NewUserRepository(cfg.DB)
	user, err := r.FindByUUID(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "User not found",
		})
	} else if err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, user)
}
