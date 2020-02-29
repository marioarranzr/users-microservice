package service

import (
	"reflect"
	"testing"

	"github.com/marioarranzr/users-microservice/domain"
	"github.com/marioarranzr/users-microservice/repository"
)

var mario = domain.User{
	FirstName: "Mario",
	LastName:  "Arranz",
	Email:     "mario@omg.lol",
}

func Test_users_Get(t *testing.T) {
	type args struct {
		u *domain.User
	}
	tests := []struct {
		name    string
		repo    []*domain.User
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "not found, empty storage",
			repo: nil,
			args: args{
				u: &domain.User{
					FirstName: "Mario",
					LastName:  "Arranz",
				},
			},
			wantErr: true,
		},
		{
			name: "found by first and last name",
			repo: []*domain.User{
				&mario,
			},
			args: args{
				u: &domain.User{
					FirstName: "Mario",
					LastName:  "Arranz",
				},
			},
			want: &mario,
		},
		{
			name: "not found (find by existing fist name and not existing last name)",
			repo: []*domain.User{
				&mario,
			},
			args: args{
				u: &domain.User{
					FirstName: "Mario",
					LastName:  "Smith",
				},
			},
			wantErr: true,
		},
		{
			name: "not found (find by not existing fist name and last name)",
			repo: []*domain.User{
				&mario,
			},
			args: args{
				u: &domain.User{
					FirstName: "Peter",
					LastName:  "Smith",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &users{
				Repo: repository.NewMemory(tt.repo),
			}
			got, err := s.Get(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("users.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("users.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_users_Post(t *testing.T) {
	type args struct {
		u *domain.User
	}
	tests := []struct {
		name    string
		repo    []*domain.User
		args    args
		wantErr bool
	}{
		{
			name: "inserted (empty database)",
			repo: nil,
			args: args{
				u: &mario,
			},
		},
		{
			name: "not inserted (user already existed)",
			repo: []*domain.User{
				&mario,
			},
			args: args{
				u: &mario,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &users{
				Repo: repository.NewMemory(tt.repo),
			}
			if err := s.Post(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("users.Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
