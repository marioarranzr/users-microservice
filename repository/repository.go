package repository

import (
	"github.com/marioarranzr/users-microservice/domain"
)

// Storage represents a place where keeping the data
type Users interface {
	Find(firstName, lastName string) (*domain.User, error)
}
