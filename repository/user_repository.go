package repository

import (
	"context"
	"log"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type UserRepository interface {
	Save(user model.User) (*model.User, error)
	FindAll() ([]model.User, error)
	DeleteAll() error
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

func (r *userRepository) Save(user model.User) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	userParams := database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	}

	dbUser, err := r.DB.CreateUser(context.Background(), userParams)
	if err != nil {
		// TODO: Return the raw sql error?
		return nil, err
	}
	createdUser := &model.User{
		ID:        dbUser.ID,
		PublicID:  dbUser.PublicID.UUID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
	}
	return createdUser, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	dbUsers, err := r.DB.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}
	users := make([]model.User, 0)
	for _, user := range dbUsers {
		u := model.User{
			ID:        user.ID,
			PublicID:  user.PublicID.UUID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteAllUsers(context.Background())
	if err != nil {
		log.Printf("error deleting all users from the db: %v", err)
		return err
	}
	return err
}
