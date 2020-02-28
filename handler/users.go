package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marioarranzr/users-microservice/domain"
	"github.com/marioarranzr/users-microservice/service"
)

type Users struct {
	Svc service.Users
}

func (u Users) Get(c echo.Context) error {
	// build user object
	user := &domain.User{}
	res, err := u.Svc.Get(user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, res)
}
