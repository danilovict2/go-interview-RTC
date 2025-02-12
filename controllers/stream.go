package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cfg *APIConfig) StreamNewToken(c echo.Context) error {
	uuid := uuidFromJWT(c)
	token, err := cfg.StreamClient.CreateToken(uuid)
	if err != nil {
		return HandleGracefully(err, c)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}