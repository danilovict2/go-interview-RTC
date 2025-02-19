package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danilovict2/go-interview-RTC/controllers"
	"github.com/danilovict2/go-interview-RTC/internal/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	api, err := controllers.NewAPIConfig()
	if err != nil {
		log.Fatal("Error creating api config:", err)
	}

	mdw := middleware.New(api.DB)

	e := echo.New()
	e.Use(mw.RequestID())
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusNotFound {
			// Instead of returning a 404, serve index.html so that Vue Router can handle the client-side route.
			if err := c.File("assets/vue/dist/index.html"); err != nil {
				c.Echo().DefaultHTTPErrorHandler(err, c)
			}
			return
		}

		c.Echo().DefaultHTTPErrorHandler(err, c)
	}

	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controllers.UserClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	e.Static("/", "assets/vue/dist")
	e.File("/", "assets/vue/dist/index.html")

	e.POST("/login", api.Login)
	e.GET("/stream/token", api.StreamNewToken, echojwt.WithConfig(jwtConfig), mdw.UUIDFromJWT)
	e.POST("/comments/store", api.CommentStore, echojwt.WithConfig(jwtConfig), mdw.UUIDFromJWT)

	u := e.Group("/users")
	u.GET("", api.UsersGet)
	u.POST("/store", api.UserStore)
	u.GET("/:uuid", api.UserGet, echojwt.WithConfig(jwtConfig), mdw.UUIDFromJWT)

	i := e.Group("/interviews", echojwt.WithConfig(jwtConfig), mdw.UUIDFromJWT)
	i.GET("", api.InterviewsGet)
	i.GET("/:stream-call-id/comments", api.InterviewGetComments)
	i.POST("/store", api.InterviewStore)
	i.PATCH("/:stream-call-id", api.InterviewPatch)

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_ADDR")))
}
