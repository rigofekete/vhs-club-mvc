package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/auth"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type UserService interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	CreateUserBatch(context.Context, []*model.User) ([]*model.User, *int32, error)
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
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	return s.repo.Save(ctx, user)
}

// TODO: Add a unit test for this
// TODO: Can we let the DB handle users that already exist, by skipping them
func (s *userService) CreateUserBatch(ctx context.Context, users []*model.User) ([]*model.User, *int32, error) {
	return s.repo.SaveBatch(ctx, users)
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
