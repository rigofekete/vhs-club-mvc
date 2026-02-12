package service_test

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTapeRespository struct {
	mock.Mock
}

func NewTapeMockRepository() *mockTapeRespository {
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

func (m *mockTapeRespository) FindByID(id int32) (*model.Tape, bool) {
	args := m.Called(id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, args.Bool(1)
}

func (m *mockTapeRespository) Update(id int32, updatedTape database.UpdateTapeParams) (*model.Tape, bool) {
	updatedTape.ID = id
	args := m.Called(id, updatedTape)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, false
}

func (m *mockTapeRespository) Delete(id int32) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *mockTapeRespository) DeleteAllTapes() bool {
	args := m.Called()
	return args.Bool(0)
}

func TestCreate_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := returnInt32(5)
	inputTape := model.Tape{
		ID:       id,
		Title:    "Sleeper",
		Director: "Woody ALlen",
		Genre:    "Comedy",
		Quantity: 1, Price: 5999.99,
	}

	createdTape := &model.Tape{
		ID:       id,
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
	mockRepo := NewTapeMockRepository()

	id := returnInt32(22)
	inputTape := model.Tape{
		ID:       id,
		Title:    "The Shining",
		Director: "",
		Genre:    "Horror",
		Quantity: 1,
		Price:    5999.99,
	}

	// mockRepo.Mock.On("Save", inputTape).Return(nil)

	svc := service.NewTapeService(mockRepo)
	tape := svc.Create(inputTape)

	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestList(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id1 := returnInt32(5)
	id2 := returnInt32(7)
	id3 := returnInt32(100)
	expectedTapes := []model.Tape{
		{
			ID:       id1,
			Title:    "Taxi Driver",
			Director: "Martin Scorcese",
			Genre:    "Thriller",
			Quantity: 1,
			Price:    5999.99,
		},
		{
			ID:       id2,
			Title:    "Amarcord",
			Director: "Federico Fellini",
			Genre:    "Drama",
			Quantity: 1,
			Price:    5999.99,
		},
		{
			ID:       id3,
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

func TestFindByID_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := returnInt32(76)
	expected := model.Tape{
		ID: id, Title: "Alien", Director: "Ridley Scott", Genre: "Horror", Quantity: 1, Price: 5999.99,
	}
	mockRepo.On("FindByID", id).Return(&expected, true)

	svc := service.NewTapeService(mockRepo)

	tape, found := svc.GetTapeByID(strconv.Itoa(int(id)))

	assert.True(t, found)
	assert.Equal(t, &expected, tape)

	mockRepo.AssertExpectations(t)
}

func TestFindByID_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := returnInt32(8)
	mockRepo.On("FindByID", id).Return(nil, false)

	svc := service.NewTapeService(mockRepo)

	tape, found := svc.GetTapeByID(strconv.Itoa(int(id)))

	assert.False(t, found)
	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := returnInt32(44)
	genre := "Difficult to label"
	partialForRepoCall := database.UpdateTapeParams{
		ID:    id,
		Genre: sql.NullString{String: genre, Valid: true},
	}
	updatedTape := model.Tape{
		ID:       id,
		Title:    "Hana-bi",
		Director: "Takeshi Kitano",
		Genre:    "Difficult to label",
		Quantity: 1,
		Price:    5999.99,
	}

	mockRepo.On("Update", id, partialForRepoCall).Return(&updatedTape, true)

	svc := service.NewTapeService(mockRepo)
	partialForSvc := model.UpdatedTape{
		Genre: &genre,
	}
	partialUpdatedTape, found := svc.Update(strconv.Itoa(int(id)), partialForSvc)

	assert.True(t, found)
	assert.Equal(t, &updatedTape, partialUpdatedTape)

	mockRepo.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := returnInt32(2001)
	partialForRepoCall := database.UpdateTapeParams{
		ID:    id,
		Title: sql.NullString{String: "Superman", Valid: true},
	}

	mockRepo.On("Update", id, partialForRepoCall).Return(nil, false)

	svc := service.NewTapeService(mockRepo)

	title := "Superman"
	partialForSvcCall := model.UpdatedTape{
		Title: &title,
	}
	updatedTape, found := svc.Update(strconv.Itoa(int(id)), partialForSvcCall)

	assert.False(t, found)
	assert.Nil(t, updatedTape)

	mockRepo.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := returnInt32(43)
	mockRepo.On("Delete", id).Return(true)

	svc := service.NewTapeService(mockRepo)
	deletedTape := svc.Delete(strconv.Itoa(int(id)))

	assert.True(t, deletedTape)

	mockRepo.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()
	id := returnInt32(20)
	mockRepo.On("Delete", id).Return(false)

	svc := service.NewTapeService(mockRepo)
	deletedTape := svc.Delete(strconv.Itoa(int(id)))

	assert.False(t, deletedTape)

	mockRepo.AssertExpectations(t)
}

func TestDeleteAllTapes(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	mockRepo.On("DeleteAllTapes").Return(true)

	svc := service.NewTapeService(mockRepo)
	deletedTape := svc.DeleteAll()

	assert.True(t, deletedTape)

	mockRepo.AssertExpectations(t)
}
