package service_test

import (
	"testing"

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

func (m *mockUserRespository) Save(user model.User) (*model.User, error) {
	args := m.Called(user)
	if t := args.Get(0); t != nil {
		return t.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) FindAll() ([]model.User, error) {
	args := m.Called()
	if users := args.Get(0); users != nil {
		return users.([]model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) DeleteAll() error {
	args := m.Called()
	return args.Error(0)
}

func Test_CreateUser_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	id := int32(14)
	inputUser := model.User{
		Name:  "Miles Davis",
		Email: "grumpy.genius@cool.com",
	}

	createdUser := &model.User{
		ID:    id,
		Name:  "Miles Davis",
		Email: "grumpy.genius@cool.com",
	}

	mockRepo.On("Save", inputUser).Return(createdUser, nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.CreateUser(inputUser)

	assert.Nil(t, err)
	assert.Equal(t, createdUser, user)

	mockRepo.AssertExpectations(t)
}

func Test_CreateUser_InvalidUser(t *testing.T) {
	mockRepo := NewUserMockRepository()

	inputUser := model.User{
		Name:  "",
		Email: "invisible@ghost.com",
	}

	// NOTE: Never called since empty strings can't be valid in the NewUser function.
	// mockRepo.Mock.On("Save", inputUser).Return(nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.CreateUser(inputUser)

	assert.Nil(t, user)
	assert.Equal(t, "invalid user fields", err.Error())
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteAllUsers(t *testing.T) {
	mockRepo := NewUserMockRepository()

	mockRepo.On("DeleteAll").Return(nil)

	svc := service.NewUserService(mockRepo)
	err := svc.DeleteAllUsers()

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
