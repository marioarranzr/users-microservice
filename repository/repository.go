package repository

import (
	"github.com/marioarranzr/users-microservice/domain"
)

// Storage represents a place where keeping the data
type Users interface {
	Find(firstName, lastName, nickname, email, country string) (*domain.User, error)
	Insert(user *domain.User) error
}
