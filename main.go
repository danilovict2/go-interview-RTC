package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "assets/vue/dist")
	e.File("/", "assets/vue/dist/index.html")

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ADDR")))
}