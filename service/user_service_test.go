package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRespository struct {
	mock.Mock
}

func NewUserMockRepository() *mockUserRespository {
	return &mockUserRespository{}
}

func (m *mockUserRespository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	args := m.Called(ctx, user)
	if u := args.Get(0); u != nil {
		return u.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) GetByID(ctx context.Context, id int32) (*model.User, error) {
	args := m.Called(ctx, id)
	if user := args.Get(0); user != nil {
		return user.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) GetByPublicID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	if user := args.Get(0); user != nil {
		return user.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) GetAll(ctx context.Context) ([]*model.User, error) {
	args := m.Called(ctx)
	if users := args.Get(0); users != nil {
		return users.([]*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) DeleteAll(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func Test_CreateUser_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	id := int32(14)
	inputUser := &model.User{
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
	}

	createdUser := &model.User{
		ID:       id,
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
	}

	ctx := context.Background()

	mockRepo.On("GetAll", ctx).Return(nil, nil)
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

	dbUsers := []*model.User{
		{
			Username: "MilesDavis",
			Email:    "one_and_only@fusion.com",
		},
		{
			Username: "Coltrane",
			Email:    "bluetrain@love.com",
		},
	}
	ctx := context.Background()

	mockRepo.On("GetAll", ctx).Return(dbUsers, nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.CreateUser(ctx, inputUser)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())

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
