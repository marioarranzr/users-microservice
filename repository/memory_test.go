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
			m := &Memory{
				Users: tt.users,
			}
			got, err := m.Find(tt.args.firstName, tt.args.lastName)
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
