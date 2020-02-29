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

func (m *Memory) Find(firstName, lastName, nickname, email, country string) (*domain.User, error) {
	for _, user := range m.Users {
		if user == nil {
			continue
		}
		if (firstName == "" || user.FirstName == firstName) &&
			(lastName == "" || user.LastName == lastName) &&
			(nickname == "" || user.Nickname == nickname) &&
			(email == "" || user.Email == email) &&
			(country == "" || user.Country == country) {

			return user, nil
		}
	}

	return nil, errors.New("not found user")
}

func (m *Memory) Insert(u *domain.User) error {
	if _, err := m.Find(u.FirstName, u.LastName, u.Nickname, u.Email, u.Country); err == nil {
		// Find NOT returning an error means the user about to insert already existed
		return errors.New("user already exists")
	}

	// Find returning an error means the user was not found
	m.Users = append(m.Users, u)
	return nil
}
