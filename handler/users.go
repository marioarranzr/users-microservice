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
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		log.Error(err)
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

func (u *Users) Put(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	res, err := u.Svc.Put(user)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, res)
}

func (u *Users) Delete(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	err := u.Svc.Delete(user)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
