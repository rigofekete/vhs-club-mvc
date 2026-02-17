package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id int32) (*model.User, error)
	GetByPublicID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	DeleteAll(ctx context.Context) error
}

type userRepository struct {
	mu sync.Mutex
	DB *database.Queries
}

func NewUserRepository() UserRepository {
	return &userRepository{
		DB: config.AppConfig.DB,
	}
}

func (r *userRepository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	userParams := database.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
	}

	dbUser, err := r.DB.CreateUser(ctx, userParams)
	if err != nil {
		return nil, err
	}
	createdUser := &model.User{
		ID:        dbUser.ID,
		PublicID:  dbUser.PublicID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
	}
	return createdUser, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (*model.User, error) {
	dbUser, err := r.DB.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", apperror.ErrUserNotFound, err)
	}
	user := &model.User{
		ID:        dbUser.ID,
		PublicID:  dbUser.PublicID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
	}
	return user, nil
}

func (r *userRepository) GetByPublicID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	dbUser, err := r.DB.GetUserFromPublicID(ctx, id)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        dbUser.ID,
		PublicID:  dbUser.PublicID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
	}

	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	dbUsers, err := r.DB.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0)
	for _, user := range dbUsers {
		u := &model.User{
			ID:        user.ID,
			PublicID:  user.PublicID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			Email:     user.Email,
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) DeleteAll(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteAllUsers(ctx)
	if err != nil {
		return err
	}
	return err
}
