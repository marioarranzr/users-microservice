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

func (m *Memory) Find(u *domain.User) ([]*domain.User, error) {
	users := []*domain.User{}
	for _, user := range m.Users {
		if user == nil {
			continue
		}
		if (u.FirstName == "" || user.FirstName == u.FirstName) &&
			(u.LastName == "" || user.LastName == u.LastName) &&
			(u.Nickname == "" || user.Nickname == u.Nickname) &&
			(u.Email == "" || user.Email == u.Email) &&
			(u.Country == "" || user.Country == u.Country) {

			user.Password = "*****"
			users = append(users, user)
		}
	}

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	return users, nil
}

func (m *Memory) Insert(u *domain.User) error {
	if _, err := m.Find(u); err == nil {
		// Find NOT returning an error means the user about to insert already existed
		return errors.New("user already exists")
	}

	// Find returning an error means the user was not found
	m.Users = append(m.Users, u)
	return nil
}

func (m *Memory) Modify(u *domain.User) (*domain.User, error) {
	var (
		user = &domain.User{
			Nickname: u.Nickname,
		}
		users []*domain.User
		err   error
	)
	if users, err = m.Find(user); err != nil {
		// Find returning an error means the user was not found
		return nil, errors.New("user not found")
	}

	// It is not possible to find more that one user with same nickname
	user = users[0]

	if u.FirstName != "" {
		user.FirstName = u.FirstName
	}
	if u.LastName != "" {
		user.LastName = u.LastName
	}
	// password is only visible when is modified
	if u.Password != "" {
		user.Password = u.Password
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Country != "" {
		user.Country = u.Country
	}

	return user, nil
}

func (m *Memory) Delete(u *domain.User) error {
	for i, user := range m.Users {
		if user == nil {
			continue
		}
		if (u.FirstName == "" || user.FirstName == u.FirstName) &&
			(u.LastName == "" || user.LastName == u.LastName) &&
			(u.Nickname == "" || user.Nickname == u.Nickname) &&
			(u.Email == "" || user.Email == u.Email) &&
			(u.Country == "" || user.Country == u.Country) {

			m.Users = append(m.Users[:i], m.Users[i+1:]...)
		}
	}
	return nil
}
