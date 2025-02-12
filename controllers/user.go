package controllers

import (
	"errors"
	"net/http"

	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

	result := cfg.DB.Create(&user)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return c.String(http.StatusBadRequest, "User with this email already exists.")
	} else if result.Error != nil {
		return HandleGracefully(err, c)
	}

	return cfg.Login(c)
}

func (cfg *APIConfig) UserGet(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid == "me" {
		uuid = uuidFromJWT(c)
	}

	user := models.User{}
	err := cfg.DB.First(&user, "uuid = ?", uuid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "User not found",
		})
	} else if err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, user)
}

func uuidFromJWT(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*UserClaims)

	return claims.UUID.String()
}
