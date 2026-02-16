package service

import (
	"context"
	"strconv"

	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type UserService interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetAllUsers(context.Context) ([]*model.User, error)
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

// TODO: Admin methods should use ID or Public ID?
func (s *userService) GetUserByID(ctx context.Context, idStr string) (*model.User, error) {
	id64, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, int32(id64))
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) DeleteAllUsers(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
