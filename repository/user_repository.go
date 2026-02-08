package repository

import (
	"context"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type UserRepository interface {
	Save(user model.User) *model.User
	FindAll() []model.User
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

func (r *userRepository) Save(user model.User) *model.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	userParams := database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	}

	dbUser, err := r.DB.CreateUser(context.Background(), userParams)
	if err != nil {
		// TODO: Should we return the err together with the object pointer?
		return nil
	}
	createdUser := &model.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
	}
	return createdUser
}

func (r *userRepository) FindAll() []model.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	dbUsers, err := r.DB.GetUsers(context.Background())
	if err != nil {
		// TODO: Should we return the err together with the object pointer?
		return nil
	}
	users := make([]model.User, 0)
	for _, user := range dbUsers {
		u := model.User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
		}
		users = append(users, u)
	}
	return users
}
