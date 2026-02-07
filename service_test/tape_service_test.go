package service_test

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
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

func (m *mockTapeRespository) FindByID(id uuid.UUID) (*model.Tape, bool) {
	args := m.Called(id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, args.Bool(1)
}

func (m *mockTapeRespository) Update(id uuid.UUID, updated database.UpdateTapeParams) (*model.Tape, bool) {
	updated.ID = id
	args := m.Called(id, updated)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, false
}

func (m *mockTapeRespository) Delete(id uuid.UUID) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestFindByID_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	id := uuid.New()
	expected := model.Tape{
		ID: id, Title: "Alien", Director: "Ridley Scott", Genre: "Horror", Quantity: 1, Price: 5999.99,
	}
	mockRepo.On("FindByID", id).Return(&expected, true)

	svc := service.NewTapeService(mockRepo)

	tape, found := svc.GetTapeByID(id.String())

	assert.True(t, found)
	assert.Equal(t, &expected, tape)

	mockRepo.AssertExpectations(t)
}

func TestFindByID_NotFound(t *testing.T) {
	mockRepo := NewMockRepository()

	id := uuid.New()
	mockRepo.On("FindByID", id).Return(nil, false)

	svc := service.NewTapeService(mockRepo)

	tape, found := svc.GetTapeByID(id.String())

	assert.False(t, found)
	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestCreate_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	id := uuid.New()
	inputTape := model.Tape{
		ID:       id,
		Title:    "Sleeper",
		Director: "Woody ALlen",
		Genre:    "Comedy",
		Quantity: 1,
		Price:    5999.99,
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
	mockRepo := NewMockRepository()

	id := uuid.New()
	inputTape := model.Tape{
		ID:       id,
		Title:    "The Shining",
		Director: "",
		Genre:    "Horror",
		Quantity: 1,
		Price:    5999.99,
	}

	mockRepo.Mock.On("Save", inputTape).Return(nil)

	svc := service.NewTapeService(mockRepo)
	tape := svc.Create(inputTape)

	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestList(t *testing.T) {
	mockRepo := NewMockRepository()

	id1 := uuid.New()
	id2 := uuid.New()
	id3 := uuid.New()
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

func TestUpdate_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	id := uuid.New()
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
	updated, found := svc.Update(id.String(), partialForSvc)

	assert.True(t, found)
	assert.Equal(t, &updatedTape, updated)

	mockRepo.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockRepo := NewMockRepository()

	id := uuid.New()
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
	updated, found := svc.Update(id.String(), partialForSvcCall)

	assert.False(t, found)
	assert.Nil(t, updated)

	mockRepo.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	mockRepo := NewMockRepository()

	id := uuid.New()
	mockRepo.On("Delete", id).Return(true)

	svc := service.NewTapeService(mockRepo)
	deleted := svc.Delete(id.String())

	assert.True(t, deleted)

	mockRepo.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	mockRepo := NewMockRepository()
	id := uuid.New()
	mockRepo.On("Delete", id).Return(false)

	svc := service.NewTapeService(mockRepo)
	deleted := svc.Delete(id.String())

	assert.False(t, deleted)

	mockRepo.AssertExpectations(t)
}
