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

func TestCreate_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	inputTape := model.Tape{
		ID:       "1",
		Title:    "Sleeper",
		Director: "Woody ALlen",
		Genre:    "Comedy",
		Quantity: 1,
		Price:    5999.99,
	}

	createdTape := &model.Tape{
		ID:       "1",
		Title:    "Sleeper",
		Director: "Woody ALlen",
		Genre:    "Comedy",
		Quantity: 1,
		Price:    5999.99,
	}

	mockRepo.On("Save", inputTape).Return(createdTape)

	svc := service.NewTapeService(mockRepo)
	tape := svc.Create(inputTape)

	assert.Equal(t, createdTape, tape)

	mockRepo.AssertExpectations(t)
}

func TestCreate_InvalidTape(t *testing.T) {
	mockRepo := NewMockRepository()

	inputTape := model.Tape{
		ID:       "1",
		Title:    "",
		Director: "Woody ALlen",
		Genre:    "Comedy",
		Quantity: 0,
		Price:    5999.99,
	}

	svc := service.NewTapeService(mockRepo)
	tape := svc.Create(inputTape)

	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestList(t *testing.T) {
	mockRepo := NewMockRepository()

	expectedTapes := []model.Tape{
		{
			ID:       "1",
			Title:    "The Terminator",
			Director: "James Cameron",
			Genre:    "Action",
			Quantity: 1,
			Price:    3999.99,
		},
		{
			ID:       "2",
			Title:    "The Predator",
			Director: "John McTiernan",
			Genre:    "Action",
			Quantity: 1,
			Price:    3999.99,
		},
		{
			ID:       "3",
			Title:    "Apocalypse Now",
			Director: "Francis Ford Coppola",
			Genre:    "Drama",
			Quantity: 1,
			Price:    5999.99,
		},
	}

	mockRepo.On("FindAll").Return(expectedTapes)

	svc := service.NewTapeService(mockRepo)
	tapes := svc.List()

	assert.Equal(t, expectedTapes, tapes)

	mockRepo.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	id := "1"
	partialUpdate := model.Tape{
		Genre: "Difficult to lable",
	}
	updatedTape := model.Tape{
		ID:       "1",
		Title:    "Hana-bi",
		Director: "Takeshi Kitano",
		Genre:    "Difficult to lable",
		Quantity: 1,
		Price:    5999.99,
	}

	mockRepo.On("Update", id, partialUpdate).Return(&updatedTape, true)

	svc := service.NewTapeService(mockRepo)
	updated, found := svc.Update(id, partialUpdate)

	assert.True(t, found)
	assert.Equal(t, &updatedTape, updated)

	mockRepo.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockRepo := NewMockRepository()

	id := "1904"
	partialUpdate := model.Tape{
		Title: "Superman",
	}

	mockRepo.On("Update", id, partialUpdate).Return(nil, false)

	svc := service.NewTapeService(mockRepo)
	updated, found := svc.Update(id, partialUpdate)

	assert.False(t, found)
	assert.Nil(t, updated)

	mockRepo.AssertExpectations(t)
}
