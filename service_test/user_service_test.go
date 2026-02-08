package service_test

import (
	"testing"

	"github.com/google/uuid"
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

func (m *mockUserRespository) Save(user model.User) *model.User {
	args := m.Called(user)
	if t := args.Get(0); t != nil {
		return t.(*model.User)
	}
	return nil
}

func (m *mockUserRespository) FindAll() []model.User {
	args := m.Called()
	if users := args.Get(0); users != nil {
		return users.([]model.User)
	}
	return nil
}

func (m *mockUserRespository) DeleteAllUsers() bool {
	args := m.Called()
	return args.Bool(0)
}

func Test_User_Create_Success(t *testing.T) {
	mockRepo := NewUserMockRepository()

	id := uuid.New()
	inputUser := model.User{
		Name:  "Miles Davis",
		Email: "grumpy.genius@cool.com",
	}

	createdUser := &model.User{
		ID:    id,
		Name:  "Miles Davis",
		Email: "grumpy.genius@cool.com",
	}

	mockRepo.On("Save", inputUser).Return(createdUser)

	svc := service.NewUserService(mockRepo)
	user := svc.Create(inputUser)

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
	user := svc.Create(inputUser)

	assert.Nil(t, user)

	mockRepo.AssertExpectations(t)
}

func TestDeleteAllUsers(t *testing.T) {
	mockRepo := NewUserMockRepository()

	mockRepo.On("DeleteAllUsers").Return(true)

	svc := service.NewUserService(mockRepo)
	deleted := svc.DeleteAll()

	assert.True(t, deleted)

	mockRepo.AssertExpectations(t)
}
