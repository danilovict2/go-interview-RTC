package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danilovict2/go-interview-RTC/controllers"
	"github.com/danilovict2/go-interview-RTC/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	api := controllers.APIConfig{
		DB: db,
	}

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controllers.UserClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	e.Static("/", "assets/vue/dist")
	e.File("/", "assets/vue/dist/index.html")

	e.POST("/login", api.Login)

	u := e.Group("/users")
	u.POST("/store", api.UserStore)
	u.GET("/:uuid", api.UserGet, echojwt.WithConfig(jwtConfig))


	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ADDR")))
}
