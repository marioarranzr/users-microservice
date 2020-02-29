package repository

import (
	"reflect"
	"testing"

	"github.com/marioarranzr/users-microservice/domain"
)

var mario = domain.User{
	FirstName: "Mario",
	LastName:  "Arranz",
	Email:     "mario@omg.lol",
}

func TestMemory_Find(t *testing.T) {
	type args struct {
		firstName string
		lastName  string
		nickname  string
		email     string
		country   string
	}
	tests := []struct {
		name    string
		users   []*domain.User
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "found by first and last name",
			users: []*domain.User{
				&mario,
			},
			args: args{
				firstName: "Mario",
				lastName:  "Arranz",
			},
			want: &mario,
		},
		{
			name:  "not found (empty database)",
			users: nil,
			args: args{
				firstName: "Mario",
				lastName:  "Smith",
			},
			wantErr: true,
		},
		{
			name: "not found (find by existing fist name and not existing last name)",
			users: []*domain.User{
				&mario,
			},
			args: args{
				firstName: "Mario",
				lastName:  "Smith",
			},
			wantErr: true,
		},
		{
			name: "not found (find by not existing fist name and last name)",
			users: []*domain.User{
				&mario,
			},
			args: args{
				firstName: "Peter",
				lastName:  "Smith",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory(tt.users)
			got, err := m.Find(tt.args.firstName, tt.args.lastName, tt.args.nickname, tt.args.email, tt.args.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("Memory.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Memory.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemory_Insert(t *testing.T) {
	type args struct {
		u *domain.User
	}
	tests := []struct {
		name    string
		users   []*domain.User
		args    args
		wantErr bool
	}{
		{
			name:  "inserted (empty database)",
			users: nil,
			args: args{
				u: &mario,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory(tt.users)
			if err := m.Insert(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Memory.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemory_Insert_Find_Update_Remove(t *testing.T) {
	type args struct {
		u *domain.User
	}
	tests := []struct {
		name  string
		users []*domain.User
		args  args
	}{
		{
			users: nil,
			args: args{
				u: &mario,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory(tt.users)
			if err := m.Insert(tt.args.u); err != nil {
				t.Errorf("Memory.Insert() error = %v", err)
			}
			got, err := m.Find(tt.args.u.FirstName, tt.args.u.LastName, tt.args.u.Nickname, tt.args.u.Email, tt.args.u.Country)
			if err != nil {
				t.Errorf("Memory.Find() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.args.u) {
				t.Errorf("Memory.Find() = %v, want %v", got, tt.args.u)
			}
			// Update
			// Remove
		})
	}
}
