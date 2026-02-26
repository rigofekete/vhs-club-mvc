package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/internal/auth"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Needed to load secret for MakeJWT calls across the service_test package
func TestMain(m *testing.M) {
	config.AppConfig = &config.Config{JWTSecret: "test-secret"}
	os.Exit(m.Run())
}

type mockUserRepository struct {
	mock.Mock
}

func NewUserMockRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	args := m.Called(ctx, user)
	if u := args.Get(0); u != nil {
		return u.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepository) SaveBatch(ctx context.Context, users []*model.User) ([]*model.User, *int32, error) {
	args := m.Called(ctx, users)
	if u := args.Get(0); u != nil {
		return u.([]*model.User), args.Get(1).(*int32), args.Error(2)
	}
	return nil, nil, args.Error(2)
}

func (m *mockUserRepository) GetByID(ctx context.Context, id int32) (*model.User, error) {
	args := m.Called(ctx, id)
	if user := args.Get(0); user != nil {
		return user.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepository) GetByPublicID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	if user := args.Get(0); user != nil {
		return user.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(ctx, username)
	if user := args.Get(0); user != nil {
		return user.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	args := m.Called(ctx)
	if users := args.Get(0); users != nil {
		return users.([]*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepository) DeleteAll(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func Test_CreateUser_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	id := int32(14)
	inputUser := &model.User{
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
		Password: "123",
	}

	createdUser := &model.User{
		ID:       id,
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
		Password: "123",
	}

	ctx := context.Background()

	mockRepo.On("Save", ctx, inputUser).Return(createdUser, nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.CreateUser(ctx, inputUser)

	assert.Nil(t, err)
	assert.Equal(t, createdUser, user)

	mockRepo.AssertExpectations(t)
}

func Test_CreateUser_Fail(t *testing.T) {
	mockRepo := NewUserMockRepository()

	inputUser := &model.User{
		Username: "MilesDavis",
		Email:    "doppelganger@ghost.com",
	}

	ctx := context.Background()

	mockRepo.On("Save", ctx, inputUser).Return(nil, apperror.ErrUserExists)

	svc := service.NewUserService(mockRepo)
	user, err := svc.CreateUser(ctx, inputUser)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())

	mockRepo.AssertExpectations(t)
}

func Test_UserLogin_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	username := "JohnCarmack"
	password := "Wolf3D"
	hashed, _ := auth.HashPassword(password)

	publicID := uuid.New()
	user := &model.User{
		PublicID:       publicID,
		Username:       username,
		Password:       password,
		HashedPassword: hashed,
	}

	foundUser := &model.User{
		PublicID:       publicID,
		Username:       username,
		Role:           "user",
		HashedPassword: hashed,
	}

	ctx := context.Background()

	mockRepo.On("GetByUsername", ctx, user.Username).Return(foundUser, nil)
	svc := service.NewUserService(mockRepo)
	loggedUser, err := svc.UserLogin(ctx, user)

	assert.Nil(t, err)
	assert.NotEqual(t, loggedUser.Token, "")

	mockRepo.AssertExpectations(t)
}

func Test_UserLogin_UserNotFound(t *testing.T) {
	mockRepo := NewUserMockRepository()

	username := "JohnRomero"
	password := "DoomGuy"
	hashed, _ := auth.HashPassword(password)

	publicID := uuid.New()
	user := &model.User{
		PublicID:       publicID,
		Username:       username,
		Password:       password,
		HashedPassword: hashed,
	}

	ctx := context.Background()

	mockRepo.On("GetByUsername", ctx, user.Username).Return(nil, apperror.ErrUserNotFound)
	svc := service.NewUserService(mockRepo)
	nullUser, err := svc.UserLogin(ctx, user)

	assert.Nil(t, nullUser)
	assert.Error(t, err)
	assert.Equal(t, err, apperror.ErrUserNotFound)

	mockRepo.AssertExpectations(t)
}

func Test_UserLogin_InvalidPW(t *testing.T) {
	mockRepo := NewUserMockRepository()

	username := "JohnRomero"
	password := "DoomGuy"
	hashed, _ := auth.HashPassword(password)

	publicID := uuid.New()
	user := &model.User{
		PublicID:       publicID,
		Username:       username,
		Password:       "1234",
		HashedPassword: hashed,
	}

	ctx := context.Background()

	mockRepo.On("GetByUsername", ctx, user.Username).Return(nil, apperror.ErrUserInvalidPW)
	svc := service.NewUserService(mockRepo)
	nullUser, err := svc.UserLogin(ctx, user)

	assert.Nil(t, nullUser)
	assert.Error(t, err)
	assert.Equal(t, err, apperror.ErrUserInvalidPW)

	mockRepo.AssertExpectations(t)
}

func Test_CreateUserBatch_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	password1 := "123"
	password2 := "magyar60"

	userBatch := []*model.User{
		{
			Username: "MilesDavis",
			Email:    "grumpy.genius@cool.com",
			Password: password1,
		},
		{
			Username: "Puskas",
			Email:    "pancho@hatharom.hu",
			Password: password2,
		},
	}

	hashPW1, _ := auth.HashPassword(password1)
	hashPW2, _ := auth.HashPassword(password2)

	savedBatch := []*model.User{
		{
			Username:       "MilesDavis",
			Email:          "grumpy.genius@cool.com",
			HashedPassword: hashPW1,
		},
		{
			Username:       "Puskas",
			Email:          "pancho@hatharom.hu",
			HashedPassword: hashPW2,
		},
	}
	ctx := context.Background()
	countArg := int32(0)

	mockRepo.On("SaveBatch", ctx, userBatch).Return(savedBatch, &countArg, nil)

	svc := service.NewUserService(mockRepo)
	users, existCount, err := svc.CreateUserBatch(ctx, userBatch)

	assert.Nil(t, err)
	assert.Equal(t, users[0].HashedPassword, hashPW1)
	assert.Equal(t, *existCount, countArg)

	mockRepo.AssertExpectations(t)
}

func Test_GetUserByID_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	id32 := int32(14)
	idUUID := uuid.New()

	returnedUser := &model.User{
		ID:       id32,
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
	}

	ctx := context.Background()

	mockRepo.On("GetByPublicID", ctx, idUUID).Return(returnedUser, nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.GetUserByID(ctx, idUUID.String())

	assert.Nil(t, err)
	assert.Equal(t, returnedUser, user)

	mockRepo.AssertExpectations(t)
}

func Test_GetUserByID_Fail(t *testing.T) {
	mockRepo := NewUserMockRepository()

	idUUID := uuid.New()

	ctx := context.Background()

	mockRepo.On("GetByPublicID", ctx, idUUID).Return(nil, apperror.ErrUserNotFound)

	svc := service.NewUserService(mockRepo)
	user, err := svc.GetUserByID(ctx, idUUID.String())

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func Test_GetAllUsers_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	ctx := context.Background()

	dbUsers := []*model.User{}

	mockRepo.On("GetAll", ctx).Return(dbUsers, nil)

	svc := service.NewUserService(mockRepo)
	users, err := svc.GetAllUsers(ctx)

	assert.Nil(t, err)
	assert.Equal(t, dbUsers, users)

	mockRepo.AssertExpectations(t)
}

func Test_DeleteAllUsers(t *testing.T) {
	mockRepo := NewUserMockRepository()

	ctx := context.Background()
	mockRepo.On("DeleteAll", ctx).Return(nil)

	svc := service.NewUserService(mockRepo)
	err := svc.DeleteAllUsers(ctx)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
