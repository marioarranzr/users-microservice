package repository

import (
	"errors"

	"github.com/marioarranzr/users-microservice/domain"
)

const (
	size = 10
)

// Memory is the In-memory database
type Memory struct {
	Users []*domain.User
}

// NewMemory creates a new In-memory database
func NewMemory(users []*domain.User) Users {
	if users == nil {
		users = make([]*domain.User, 0, size)
	}
	return &Memory{
		Users: users,
	}
}

func (m *Memory) Find(firstName, lastName string) (*domain.User, error) {
	for _, user := range m.Users {
		if user == nil {
			continue
		}
		if (firstName == "" || user.FirstName == firstName) &&
			(lastName == "" || user.LastName == lastName) {
			return user, nil
		}
	}

	return nil, errors.New("not found user")
}
