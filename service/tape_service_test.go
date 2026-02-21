package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTapeRepository struct {
	mock.Mock
}

func NewTapeMockRepository() *mockTapeRepository {
	return &mockTapeRepository{}
}

func (m *mockTapeRepository) Save(ctx context.Context, tape *model.Tape) (*model.Tape, error) {
	args := m.Called(ctx, tape)
	if t := args.Get(0); t != nil {
		return t.(*model.Tape), args.Error(1)
	}
	return nil, errors.New("invalid mock tape fields")
}

func (m *mockTapeRepository) GetAll(ctx context.Context) ([]*model.Tape, error) {
	args := m.Called(ctx)
	if tapes := args.Get(0); tapes != nil {
		return tapes.([]*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRepository) GetByID(ctx context.Context, id int32) (*model.Tape, error) {
	args := m.Called(ctx, id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRepository) GetByPublicID(ctx context.Context, id uuid.UUID) (*model.Tape, error) {
	args := m.Called(ctx, id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRepository) Update(ctx context.Context, updateTape *model.UpdateTape) (*model.Tape, error) {
	args := m.Called(ctx, updateTape)
	if updTape := args.Get(0); updTape != nil {
		return updTape.(*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRepository) Delete(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTapeRepository) DeleteAll(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func Test_CreateTape_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := int32(5)
	inputTape := &model.Tape{
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

	ctx := context.Background()

	mockRepo.On("GetAll", ctx).Return(nil, nil)
	mockRepo.On("Save", ctx, inputTape).Return(createdTape, nil)

	svc := service.NewTapeService(mockRepo)
	tape, err := svc.CreateTape(ctx, inputTape)

	assert.Equal(t, createdTape, tape)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_CreateTape_Fail(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	inputTape := &model.Tape{
		Title:    "The Shining",
		Director: "Stanley Kubrick",
		Genre:    "Horror",
		Quantity: 1,
		Price:    5999.99,
	}

	dbTapes := []*model.Tape{
		{
			Title:    "The Shining",
			Director: "Stanley Kubrick",
			Genre:    "Horror",
			Quantity: 1,
			Price:    5999.99,
		},
	}

	ctx := context.Background()

	mockRepo.On("GetAll", ctx).Return(dbTapes, nil)

	svc := service.NewTapeService(mockRepo)
	tape, err := svc.CreateTape(ctx, inputTape)

	assert.Nil(t, tape)
	assert.Error(t, err)
	assert.Equal(t, "tape already exists", err.Error())

	mockRepo.AssertExpectations(t)
}

func Test_GetAllTapes(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id1 := int32(5)
	id2 := int32(7)
	id3 := int32(100)
	expectedTapes := []*model.Tape{
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

	ctx := context.Background()

	mockRepo.On("GetAll", ctx).Return(expectedTapes, nil)

	svc := service.NewTapeService(mockRepo)
	tapes, err := svc.GetAllTapes(ctx)

	assert.Equal(t, expectedTapes, tapes)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_GetTapeByID_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	id32 := int32(76)
	idUUID := uuid.New()

	returnedTape := &model.Tape{
		ID:       id32,
		PublicID: idUUID,
		Title:    "Alien",
		Director: "Ridley Scott",
		Genre:    "Horror",
		Quantity: 1,
		Price:    5999.99,
	}

	mockRepo.On("GetByPublicID", ctx, idUUID).Return(returnedTape, nil)

	svc := service.NewTapeService(mockRepo)

	tape, err := svc.GetTapeByID(ctx, idUUID.String())

	assert.Nil(t, err)
	assert.Equal(t, returnedTape, tape)

	mockRepo.AssertExpectations(t)
}

func TestGetTapeByID_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	idUUID := uuid.New()
	mockRepo.On("GetByPublicID", ctx, idUUID).Return(nil, apperror.ErrTapeNotFound)

	svc := service.NewTapeService(mockRepo)

	tape, err := svc.GetTapeByID(ctx, idUUID.String())

	assert.Error(t, err)
	assert.Equal(t, "tape not found", err.Error())
	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateTape_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id32 := int32(44)
	idUUID := uuid.New()
	genre := "Difficult to label"
	partialForRepoCall := &model.UpdateTape{
		ID:    id32,
		Genre: &genre,
	}

	originalTape := &model.Tape{
		ID:       id32,
		Title:    "Hana-bi",
		Director: "Takeshi Kitano",
		Genre:    "Drama",
		Quantity: 1,
		Price:    5999.99,
	}

	updatedTape := &model.Tape{
		ID:       id32,
		Title:    "Hana-bi",
		Director: "Takeshi Kitano",
		Genre:    "Difficult to label",
		Quantity: 1,
		Price:    5999.99,
	}

	ctx := context.Background()
	mockRepo.On("GetByPublicID", ctx, idUUID).Return(originalTape, nil)
	mockRepo.On("Update", ctx, partialForRepoCall).Return(updatedTape, nil)
	// mockRepo.On("Exists", ctx, id).Return(true, nil)

	svc := service.NewTapeService(mockRepo)
	partialForSvc := &model.UpdateTape{
		Genre: &genre,
	}
	partialUpdatedTape, err := svc.UpdateTape(ctx, idUUID.String(), partialForSvc)

	assert.Nil(t, err)
	assert.Equal(t, updatedTape, partialUpdatedTape)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateTape_Fail(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	idUUID := uuid.New()
	title := "Superman"
	ctx := context.Background()
	mockRepo.On("GetByPublicID", ctx, idUUID).Return(nil, apperror.ErrTapeNotFound)

	svc := service.NewTapeService(mockRepo)

	partialForSvcCall := &model.UpdateTape{
		Title: &title,
	}
	updatedTape, err := svc.UpdateTape(ctx, idUUID.String(), partialForSvcCall)

	assert.Error(t, err)
	assert.Equal(t, "tape not found", err.Error())
	assert.Nil(t, updatedTape)

	mockRepo.AssertExpectations(t)
}

func Test_DeleteTape_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	idUUID := uuid.New()
	id32 := int32(14)

	returnedTape := &model.Tape{
		ID:       id32,
		PublicID: idUUID,
		Title:    "Alien",
		Director: "Ridley Scott",
		Genre:    "Horror",
		Quantity: 5,
		Price:    5999.99,
	}
	ctx := context.Background()
	mockRepo.On("GetByPublicID", ctx, idUUID).Return(returnedTape, nil)
	mockRepo.On("Delete", ctx, id32).Return(nil)

	svc := service.NewTapeService(mockRepo)
	err := svc.DeleteTape(ctx, idUUID.String())

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_DeleteTape_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	idUUID := uuid.New()
	mockRepo.On("GetByPublicID", ctx, idUUID).Return(nil, apperror.ErrTapeNotFound)

	svc := service.NewTapeService(mockRepo)
	err := svc.DeleteTape(ctx, idUUID.String())

	assert.Error(t, err)
	assert.Equal(t, "tape not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func Test_DeleteAllTapes(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	mockRepo.On("DeleteAll", ctx).Return(nil)

	svc := service.NewTapeService(mockRepo)
	err := svc.DeleteAllTapes(ctx)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
