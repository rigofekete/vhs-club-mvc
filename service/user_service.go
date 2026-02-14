package service

import (
	"errors"

	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type UserService interface {
	Create(model.User) (*model.User, error)
	List() ([]model.User, error)
	DeleteAll() error
}

type userService struct {
	repo repository.UserRepository
}

func validUserFields(user model.User) bool {
	return user.Name != "" && user.Email != ""
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) Create(user model.User) (*model.User, error) {
	if !validUserFields(user) {
		return nil, errors.New("invalid user fields")
	}
	return s.repo.Save(user)
}

func (s *userService) List() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) DeleteAll() error {
	return s.repo.DeleteAllUsers()
}
