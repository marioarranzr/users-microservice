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
	tests := []struct {
		name    string
		users   []*domain.User
		args    *domain.User
		want    *domain.User
		wantErr bool
	}{
		{
			name: "found by first and last name",
			users: []*domain.User{
				&mario,
			},
			args: &domain.User{
				FirstName: "Mario",
				LastName:  "Arranz",
			},
			want: &mario,
		},
		{
			name:  "not found (empty database)",
			users: nil,
			args: &domain.User{
				FirstName: "Mario",
				LastName:  "Smith",
			},
			wantErr: true,
		},
		{
			name: "not found (find by existing fist name and not existing last name)",
			users: []*domain.User{
				&mario,
			},
			args: &domain.User{
				FirstName: "Mario",
				LastName:  "Smith",
			},
			wantErr: true,
		},
		{
			name: "not found (find by not existing fist name and last name)",
			users: []*domain.User{
				&mario,
			},
			args: &domain.User{
				FirstName: "Peter",
				LastName:  "Smith",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMemory(tt.users)
			gotList, err := m.Find(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Memory.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				got := gotList[0]
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Memory.Find() = %v, want %v", got, tt.want)
				}
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
		{
			name: "not inserted (user already exist)",
			users: []*domain.User{
				&mario,
			}, args: args{
				u: &mario,
			},
			wantErr: true,
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

func TestMemory_Insert_Find_Update_Delete(t *testing.T) {
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
			gotList, err := m.Find(tt.args.u)
			if err != nil {
				t.Errorf("Memory.Find() error = %v", err)
				return
			}
			got := gotList[0]
			if !reflect.DeepEqual(got, tt.args.u) {
				t.Errorf("Memory.Find() = %v, want %v", got, tt.args.u)
			}
			tt.args.u.LastName = "Smith"
			got, err = m.Modify(tt.args.u)
			if err != nil {
				t.Errorf("Memory.Modify() error = %v", err)
				return
			}
			if got.LastName != tt.args.u.LastName {
				t.Errorf("Memory.Modify() = %v, want %v", got.LastName, tt.args.u.LastName)
			}
			err = m.Delete(tt.args.u)
			if err != nil {
				t.Errorf("Memory.Delete() error = %v", err)
				return
			}
		})
	}
}
