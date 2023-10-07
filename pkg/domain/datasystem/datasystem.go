package datasystem

import (
	"fmt"

	"github.com/LCRERGO/drucssGO/pkg/domain/entity"
	"github.com/google/uuid"
)

type Datasystem interface {
	CreateUser(name string)
	ReadAllUsers() []entity.User
	ReadUser(id uuid.UUID) (*entity.User, error)
	UpdateUser(id uuid.UUID, name string) error
	DeleteUser(id uuid.UUID) error
}

type datasystem struct {
	users []entity.User
}

func NewDatasystem() Datasystem {
	return &datasystem{
		users: make([]entity.User, 0),
	}
}

func (d *datasystem) CreateUser(name string) {
	d.users = append(d.users, entity.User{
		ID:   uuid.New(),
		Name: name,
	})
}

func (d *datasystem) ReadAllUsers() []entity.User {
	return d.users
}

func (d *datasystem) ReadUser(id uuid.UUID) (*entity.User, error) {
	for i, user := range d.users {
		if user.ID == id {
			return &d.users[i], nil
		}
	}

	return nil, fmt.Errorf("could not find user %v", id)
}

func (d *datasystem) UpdateUser(id uuid.UUID, name string) error {
	for i, user := range d.users {
		if user.ID == id {
			d.users[i].Name = name
			return nil
		}
	}

	return fmt.Errorf("could not update user %v", id)
}

func (d *datasystem) DeleteUser(id uuid.UUID) error {
	for i, user := range d.users {
		if user.ID == id {
			d.users = append(d.users[:i], d.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("could not delete user %v", id)
}
