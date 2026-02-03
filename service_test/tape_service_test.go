package service_test

import (
	"testing"

	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTapeRespository struct {
	mock.Mock
}

func NewMockRepository() *mockTapeRespository {
	return &mockTapeRespository{}
}

func (m *mockTapeRespository) Save(tape model.Tape) *model.Tape {
	args := m.Called(tape)
	if t := args.Get(0); t != nil {
		return t.(*model.Tape)
	}
	return nil
}

func (m *mockTapeRespository) FindAll() []model.Tape {
	args := m.Called()
	if tapes := args.Get(0); tapes != nil {
		return tapes.([]model.Tape)
	}
	return nil
}

func (m *mockTapeRespository) FindByID(id string) (*model.Tape, bool) {
	args := m.Called(id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, args.Bool(1)
}

func (m *mockTapeRespository) Update(id string, updated model.Tape) (*model.Tape, bool) {
	args := m.Called(id, updated)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, false
}

func (m *mockTapeRespository) Delete(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestFindByID_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	expected := model.Tape{
		ID: "1", Title: "Alien", Director: "Ridley Scott", Genre: "Horror", Quantity: 1, Price: 5999.99,
	}
	mockRepo.On("FindByID", "1").Return(&expected, true)

	svc := service.NewTapeService(mockRepo)

	tape, found := svc.GetTapeByID("1")

	assert.True(t, found)
	assert.Equal(t, &expected, tape)

	mockRepo.AssertExpectations(t)
}

func TestFindByID_NotFound(t *testing.T) {
	mockRepo := NewMockRepository()

	mockRepo.On("FindByID", "22").Return(nil, false)

	svc := service.NewTapeService(mockRepo)

	tape, found := svc.GetTapeByID("22")

	assert.False(t, found)
	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := NewMockRepository()

	expectation := model.Tape{
		ID:       "1",
		Title:    "Sleeper",
		Director: "Woody ALlen",
		Genre:    "Comedy",
		Quantity: 1,
		Price:    5999.99,
	}

	mockRepo.On("Save", expectation).Return(&expectation)

	svc := service.NewTapeService(mockRepo)
	tape := svc.Create(expectation)

	assert.Equal(t, &expectation, tape)

	mockRepo.AssertExpectations(t)
}

// func Test

// func TestFindByID_Success(t *testing.T) {
// 	mockRepo := NewMockRepository()
//
// 	expected := model.Tape{
// 		ID: "1", Title: "Alien", Director: "Ridley Scott", Genre: "Horror", Quantity: 1, Price: 5999.99,
// 	}
// 	mockRepo.On("FindByID", "1").Return(&expected, true)
//
// 	svc := service.NewTapeService(mockRepo)
//
// 	tape, found := svc.GetTapeByID("1")
//
// 	assert.True(t, found)
// 	assert.Equal(t, &expected, tape)
//
// 	mockRepo.AssertExpectations(t)
// }
//
