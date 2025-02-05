package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGracefully(err error, c echo.Context) error {
	c.Logger().Error(err)
	return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
}
