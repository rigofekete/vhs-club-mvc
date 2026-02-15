package servicetest

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

func (m *mockUserRespository) Save(user *model.User) (*model.User, error) {
	args := m.Called(user)
	if t := args.Get(0); t != nil {
		return t.(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRespository) FindAll() ([]*model.User, error) {
	args := m.Called()
	if users := args.Get(0); users != nil {
		return users.([]*model.User), args.Error(1)
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
	inputUser := &model.User{
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
	}

	createdUser := &model.User{
		ID:       id,
		Username: "MilesDavis",
		Email:    "grumpy.genius@cool.com",
	}

	mockRepo.On("Save", inputUser).Return(createdUser, nil)

	svc := service.NewUserService(mockRepo)
	user, err := svc.CreateUser(inputUser)

	assert.Nil(t, err)
	assert.Equal(t, createdUser, user)

	mockRepo.AssertExpectations(t)
}

// TODO: Invalid user test needs to comply with the scope of the service layer validation. User exists in the DB, etc.
// TODO: Needs to be refactored
// func Test_CreateUser_InvalidUser(t *testing.T) {
// 	mockRepo := NewUserMockRepository()
//
// 	inputUser := &model.User{
// 		Username: "",
// 		Email:    "invisible@ghost.com",
// 	}
//
// 	mockRepo.Mock.On("Save", inputUser).Return(nil, errors.New("invalid user fields"))
//
// 	svc := service.NewUserService(mockRepo)
// 	user, err := svc.CreateUser(inputUser)
//
// 	assert.Nil(t, user)
// 	assert.Equal(t, "invalid user fields", err.Error())
// 	assert.Error(t, err)
//
// 	mockRepo.AssertExpectations(t)
// }

func TestDeleteAllUsers(t *testing.T) {
	mockRepo := NewUserMockRepository()

	mockRepo.On("DeleteAll").Return(nil)

	svc := service.NewUserService(mockRepo)
	err := svc.DeleteAllUsers()

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
