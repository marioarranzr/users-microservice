package service

import (
	"errors"

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

func (s *users) Get(u *domain.User) ([]*domain.User, error) {
	return s.Repo.Find(u)
}

func (s *users) Post(u *domain.User) error {
	if u.Nickname == "" {
		return errors.New("nickname is mandatory")
	}
	return s.Repo.Insert(u)
}

func (s *users) Put(u *domain.User) (*domain.User, error) {
	if u.Nickname == "" {
		return nil, errors.New("nickname is mandatory")
	}
	return s.Repo.Modify(u)
}

func (s *users) Delete(u *domain.User) error {
	return s.Repo.Delete(u)
}
