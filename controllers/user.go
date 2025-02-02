package controllers

import (
	"errors"
	"net/http"

	"github.com/danilovict2/go-interview-RTC/internal/database"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserStore(c echo.Context) error {
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
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
	}

	db, err := database.NewConnection()
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
	}

	result := db.Create(&user)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return c.String(http.StatusBadRequest, "User with this email already exists.")
	} else if result.Error != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
	}

	return Login(c)
}
