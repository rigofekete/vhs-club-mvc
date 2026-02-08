package service

import (
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type UserService interface {
	Create(model.User) *model.User
	List() []model.User
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) Create(user model.User) *model.User {
	return s.repo.Save(user)
}

func (s *userService) List() []model.User {
	return s.repo.FindAll()
}
