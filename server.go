package main

import (
	"net/http"

	"encompass/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/login", api.LoginHandler)
	e.GET("/callback", api.CallbackHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
