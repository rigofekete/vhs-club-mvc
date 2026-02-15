package repository

import (
	"context"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type UserRepository interface {
	Save(context.Context, *model.User) (*model.User, error)
	FindAll(context.Context) ([]*model.User, error)
	DeleteAll(context.Context) error
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
		PublicID:  dbUser.PublicID.UUID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
	}
	return createdUser, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
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
			PublicID:  user.PublicID.UUID,
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
