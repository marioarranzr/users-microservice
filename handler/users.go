package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/marioarranzr/users-microservice/domain"
	"github.com/marioarranzr/users-microservice/service"
)

type Users struct {
	Svc service.Users
}

func (u *Users) Get(c echo.Context) error {
	// build user object
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	res, err := u.Svc.Get(user)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, res)
}

func (u *Users) Post(c echo.Context) error {
	// build user object
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	err := u.Svc.Post(user)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusConflict)
	}

	return c.JSON(http.StatusCreated, user)
}
