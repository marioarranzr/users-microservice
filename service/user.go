package service

import (
	"github.com/marioarranzr/users-microservice/domain"
	"github.com/marioarranzr/users-microservice/repository"
)

type users struct {
	Repo repository.Users
}

// NewService creates an adding service with the necessary dependencies
func New(r repository.Users) Users {
	return &users{
		Repo: r,
	}
}

func (s *users) Get(u *domain.User) (*domain.User, error) {
	return s.Repo.Find(u.FirstName, u.LastName)
}
