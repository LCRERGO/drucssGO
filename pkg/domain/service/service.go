package service

import (
	"github.com/LCRERGO/drucssGO/pkg/domain/datasystem"
	"github.com/LCRERGO/drucssGO/pkg/domain/entity"
)

type Service interface {
	CreateUser(entity.User) error
	ReadAllUsers() ([]entity.User, error)
	ReadUser(entity.User) (*entity.User, error)
	UpdateUser(entity.User) error
	DeleteUser(entity.User) error
}

type service struct {
	datasystem datasystem.Datasystem
}

func NewService() Service {
	return &service{
		datasystem: datasystem.NewDatasystem(),
	}
}

func (s *service) CreateUser(inUser entity.User) error {
	s.datasystem.CreateUser(inUser.Name)

	return nil
}

func (s *service) ReadAllUsers() ([]entity.User, error) {
	return s.datasystem.ReadAllUsers(), nil
}

func (s *service) ReadUser(inUser entity.User) (*entity.User, error) {
	user, err := s.datasystem.ReadUser(inUser.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) UpdateUser(inUser entity.User) error {
	err := s.datasystem.UpdateUser(inUser.ID, inUser.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUser(inUser entity.User) error {
	err := s.datasystem.DeleteUser(inUser.ID)
	if err != nil {
		return err
	}

	return nil
}
