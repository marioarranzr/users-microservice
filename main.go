package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/marioarranzr/users-microservice/handler"
	"github.com/marioarranzr/users-microservice/repository"
	"github.com/marioarranzr/users-microservice/service"
)

const (
	port = ":9091"
)

func main() {
	var (
		repo repository.Users
		svc  service.Users
	)
	e := echo.New()

	// Logger
	e.Use(middleware.Logger())
	log.SetLevel(log.INFO)

	// In-memory database
	repo = repository.NewMemory(nil)
	svc = service.New(repo)

	// Handlers
	loadUsers(e, svc)

	e.Logger.Fatal(e.Start(port))
}

func loadUsers(e *echo.Echo, svc service.Users) {
	u := &handler.Users{
		Svc: svc,
	}
	e.GET("/users", u.Get)
}
