package middleware

import (
	"errors"
	"net/http"

	"github.com/danilovict2/go-interview-RTC/controllers"
	"github.com/danilovict2/go-interview-RTC/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (mw *Middleware) UserFromJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := uuidFromJWT(c)
		r := repository.NewUserRepository(mw.DB)
		user, err := r.FindByUUID(uuid)
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "User not found",
			})
		} else if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
		}

		c.Set("authUser", user)
		return next(c)
	}
}

func (mw *Middleware) UUIDFromJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := uuidFromJWT(c)
		c.Set("uuid", uuid)
		return next(c)
	}
}

func uuidFromJWT(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*controllers.UserClaims)

	return claims.UUID.String()
}
