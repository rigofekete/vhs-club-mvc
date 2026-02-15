package service

import (
	"context"

	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type UserService interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	ListUsers(context.Context) ([]*model.User, error)
	DeleteAllUsers(context.Context) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.Save(ctx, user)
}

func (s *userService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *userService) DeleteAllUsers(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
