package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIConfig struct {
	DB *gorm.DB
}

func HandleGracefully(err error, c echo.Context) error {
	c.Logger().Error(err)
	return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
}
