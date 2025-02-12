package controllers

import (
	"net/http"
	"os"

	"github.com/GetStream/getstream-go"
	"github.com/danilovict2/go-interview-RTC/internal/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type APIConfig struct {
	DB           *gorm.DB
	StreamClient *getstream.Stream
}

func NewAPIConfig() (*APIConfig, error) {
	db, err := database.NewConnection()
	if err != nil {
		return &APIConfig{}, err
	}

	streamClient, err := getstream.NewClient(os.Getenv("STREAM_API_KEY"), os.Getenv("STREAM_API_SECRET"))
	if err != nil {
		return &APIConfig{}, err
	}

	return &APIConfig{
		DB:           db,
		StreamClient: streamClient,
	}, nil
}

func HandleGracefully(err error, c echo.Context) error {
	c.Logger().Error(err)
	return c.String(http.StatusInternalServerError, "Whoops, something went wrong!")
}
