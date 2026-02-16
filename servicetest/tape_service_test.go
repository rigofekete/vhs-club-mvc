package servicetest

import (
	"context"
	"errors"
	"strconv"
	"testing"

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

func (m *mockTapeRespository) Save(ctx context.Context, tape *model.Tape) (*model.Tape, error) {
	args := m.Called(ctx, tape)
	if t := args.Get(0); t != nil {
		return t.(*model.Tape), args.Error(1)
	}
	return nil, errors.New("invalid mock tape fields")
}

func (m *mockTapeRespository) FindAll(ctx context.Context) ([]*model.Tape, error) {
	args := m.Called(ctx)
	if tapes := args.Get(0); tapes != nil {
		return tapes.([]*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRespository) FindByID(ctx context.Context, id int32) (*model.Tape, error) {
	args := m.Called(ctx, id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRespository) Update(ctx context.Context, updateTape *model.UpdateTape) (*model.Tape, error) {
	args := m.Called(ctx, updateTape)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockTapeRespository) Delete(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockTapeRespository) DeleteAll(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// TODO: Refactor UpdateTape tests to handle this
func (m *mockTapeRespository) Exists(ctx context.Context, id int32) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func TestCreateTape_Success(t *testing.T) {
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

	mockRepo.On("Save", ctx, inputTape).Return(createdTape, nil)

	svc := service.NewTapeService(mockRepo)
	tape, err := svc.CreateTape(ctx, inputTape)

	assert.Equal(t, createdTape, tape)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

// func TestCreateTape_InvalidTape(t *testing.T) {
// 	mockRepo := NewTapeMockRepository()
//
// 	id := int32(22)
// 	inputTape := &model.Tape{
// 		ID:       id,
// 		Title:    "The Shining",
// 		Director: "",
// 		Genre:    "Horror",
// 		Quantity: 1,
// 		Price:    5999.99,
// 	}
//
// 	ctx := context.Background()
//
// 	svc := service.NewTapeService(mockRepo)
// 	tape, err := svc.CreateTape(ctx, inputTape)
//
// 	assert.Nil(t, tape)
// 	assert.Error(t, err)
// 	assert.Equal(t, "invalid tape fields", err.Error())
//
// 	mockRepo.AssertExpectations(t)
// }

func TestListTapes(t *testing.T) {
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

	mockRepo.On("FindAll", ctx).Return(expectedTapes, nil)

	svc := service.NewTapeService(mockRepo)
	tapes, err := svc.ListTapes(ctx)

	assert.Equal(t, expectedTapes, tapes)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetTapeByID_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	id := int32(76)
	expected := model.Tape{
		ID: id, Title: "Alien", Director: "Ridley Scott", Genre: "Horror", Quantity: 1, Price: 5999.99,
	}
	mockRepo.On("FindByID", ctx, id).Return(&expected, nil)

	svc := service.NewTapeService(mockRepo)

	tape, err := svc.GetTapeByID(ctx, strconv.Itoa(int(id)))

	assert.Nil(t, err)
	assert.Equal(t, &expected, tape)

	mockRepo.AssertExpectations(t)
}

func TestGetTapeByID_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	id := int32(8)
	mockRepo.On("FindByID", ctx, id).Return(nil, errors.New("mock tape not found"))

	svc := service.NewTapeService(mockRepo)

	tape, err := svc.GetTapeByID(ctx, strconv.Itoa(int(id)))

	assert.Error(t, err)
	assert.Equal(t, "mock tape not found", err.Error())
	assert.Nil(t, tape)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTape_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := int32(44)
	genre := "Difficult to label"
	partialForRepoCall := &model.UpdateTape{
		ID:    id,
		Genre: &genre,
	}
	updatedTape := model.Tape{
		ID:       id,
		Title:    "Hana-bi",
		Director: "Takeshi Kitano",
		Genre:    "Difficult to label",
		Quantity: 1,
		Price:    5999.99,
	}

	ctx := context.Background()
	mockRepo.On("Update", ctx, partialForRepoCall).Return(&updatedTape, nil)
	mockRepo.On("Exists", ctx, id).Return(true, nil)

	svc := service.NewTapeService(mockRepo)
	partialForSvc := &model.UpdateTape{
		Genre: &genre,
	}
	partialUpdatedTape, err := svc.UpdateTape(ctx, strconv.Itoa(int(id)), partialForSvc)

	assert.Nil(t, err)
	assert.Equal(t, &updatedTape, partialUpdatedTape)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTape_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	id := int32(2001)
	title := "Superman"
	ctx := context.Background()
	mockRepo.On("Exists", ctx, id).Return(false, errors.New("mock tape not found"))

	svc := service.NewTapeService(mockRepo)

	partialForSvcCall := &model.UpdateTape{
		Title: &title,
	}
	updatedTape, err := svc.UpdateTape(ctx, strconv.Itoa(int(id)), partialForSvcCall)

	assert.Error(t, err)
	assert.Equal(t, "mock tape not found", err.Error())
	assert.Nil(t, updatedTape)

	mockRepo.AssertExpectations(t)
}

func TestDeleteTape_Success(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	id := int32(43)
	mockRepo.On("Delete", ctx, id).Return(nil)

	svc := service.NewTapeService(mockRepo)
	err := svc.DeleteTape(ctx, strconv.Itoa(int(id)))

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteTape_NotFound(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	id := int32(20)
	mockRepo.On("Delete", ctx, id).Return(errors.New("mock tape not found"))

	svc := service.NewTapeService(mockRepo)
	err := svc.DeleteTape(ctx, strconv.Itoa(int(id)))

	assert.Error(t, err)
	assert.Equal(t, "mock tape not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDeleteAllTapes(t *testing.T) {
	mockRepo := NewTapeMockRepository()

	ctx := context.Background()
	mockRepo.On("DeleteAll", ctx).Return(nil)

	svc := service.NewTapeService(mockRepo)
	err := svc.DeleteAllTapes(ctx)

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
