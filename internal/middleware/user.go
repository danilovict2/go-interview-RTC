package middleware

import (
	"github.com/danilovict2/go-interview-RTC/controllers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

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
