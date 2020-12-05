package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func start() {
	e := echo.New()
	e.GET("/users/get", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":5555"))
}
