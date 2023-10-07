package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/login", LoginHandler)
	e.GET("/callback", CallbackHandler)
	return e
}
