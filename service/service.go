package service

import (
	"github.com/marioarranzr/users-microservice/domain"
)

type Users interface {
	Get(user *domain.User) (*domain.User, error)
	Post(user *domain.User) error
}
