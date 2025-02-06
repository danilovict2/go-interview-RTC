package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/danilovict2/go-interview-RTC/internal/database"
	"github.com/danilovict2/go-interview-RTC/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	UUID uuid.UUID
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	db, err := database.NewConnection()
	if err != nil {
		return HandleGracefully(err, c)
	}

	user := &models.User{}
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return c.String(http.StatusUnauthorized, "User with this email doesn't exist!")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return c.String(http.StatusUnauthorized, "The Email or Password is Incorrect. Try again.")
	}

	claims := &UserClaims{
		user.UUID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return HandleGracefully(err, c)
	}

	expires, err := token.Claims.GetExpirationTime()
	if err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   t,
		"expires": expires.Time.UTC().Format(http.TimeFormat),
	})
}
