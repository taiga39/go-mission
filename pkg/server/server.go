package server

import (
	"go-mission/pkg/controller"

	"github.com/labstack/echo"
)

func Start() {
	e := echo.New()
	e.POST("/users/create", controller.CreateUser)
	e.Logger.Fatal(e.Start(":5555"))
}
