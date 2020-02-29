package repository

import (
	"github.com/marioarranzr/users-microservice/domain"
)

// Users represents the users repository interface
type Users interface {
	Find(user *domain.User) ([]*domain.User, error)
	Insert(user *domain.User) error
	Modify(user *domain.User) (*domain.User, error)
	Delete(user *domain.User) error
}
