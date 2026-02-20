package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
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
	dbUsers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, dbUser := range dbUsers {
		if dbUser.Username == user.Username {
			return nil, apperror.ErrUserExists
		}
	}

	return s.repo.Save(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByPublicID(ctx, idUUID)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) DeleteAllUsers(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
