package service_test

import (
	"errors"
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
		return t.(*model.User), nil
	}
	return nil, errors.New("error occured")
}

func (m *mockUserRespository) FindAll() ([]model.User, error) {
	args := m.Called()
	if users := args.Get(0); users != nil {
		return users.([]model.User), nil
	}
	return nil, errors.New("error occured")
}

func (m *mockUserRespository) DeleteAllUsers() error {
	args := m.Called()
	return args.Error(0)
}

func Test_User_Create_Success(t *testing.T) {
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
	user, err := svc.Create(inputUser)

	assert.Nil(t, err)
	assert.Equal(t, createdUser, user)

	mockRepo.AssertExpectations(t)
}

func Test_Create_InvalidUser(t *testing.T) {
	mockRepo := NewUserMockRepository()

	inputUser := model.User{
		Name:  "",
		Email: "invisible@ghost.com",
	}

	// NOTE: Never called since empty strings can't be valid in the Create function.
	// mockRepo.Mock.On("Save", inputUser).Return(nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.Create(inputUser)

	assert.Nil(t, user)
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteAllUsers(t *testing.T) {
	mockRepo := NewUserMockRepository()

	mockRepo.On("DeleteAllUsers").Return(nil)

	svc := service.NewUserService(mockRepo)
	err := svc.DeleteAll()

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
