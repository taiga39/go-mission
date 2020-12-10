package controller

import (
	"go-mission/pkg/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string
}

func CreateUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	usecase.CreateUser(u.Name)
	return c.String(http.StatusOK, "A successful response.")
}
