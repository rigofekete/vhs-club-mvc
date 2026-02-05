package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rigofekete/vhs-club-mvc/handler"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTapeService struct {
	mock.Mock
}

func NewMockTapeService() *mockTapeService {
	return &mockTapeService{}
}

func (m *mockTapeService) GetTapeByID(id string) (*model.Tape, bool) {
	args := m.Called(id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), true
	}
	return nil, false
}

func (m *mockTapeService) Create(tape model.Tape) *model.Tape {
	args := m.Called(tape)
	if t := args.Get(0); t != nil {
		return t.(*model.Tape)
	}
	return nil
}

func (m *mockTapeService) Update(id string, tape model.Tape) (*model.Tape, bool) {
	args := m.Called(id, tape)
	if t := args.Get(0); t != nil {
		return t.(*model.Tape), true
	}
	return nil, false
}

func (m *mockTapeService) Delete(id string) bool {
	args := m.Called(id)
	return false != args.Bool(0)
}

func (m *mockTapeService) List() []model.Tape {
	args := m.Called()
	if list := args.Get(0); list != nil {
		return list.([]model.Tape)
	}
	return nil
}

func TestGetTapeByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := NewMockTapeService()
	expected := &model.Tape{
		ID: "1", Title: "Taxi Driver",
		Director: "Martin Scorcese", Genre: "Thriller",
		Quantity: 1, Price: 5999.99,
	}
	mockSvc.On("GetTapeByID", "1").Return(expected, true)

	h := handler.NewTapeHandler(mockSvc)

	r := gin.Default()
	r.GET("/tapes/:id", h.GetTapeByID)

	req, _ := http.NewRequest("GET", "/tapes/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectation := `
	{"id": "1", "title": "Taxi Driver", 
		"director": "Martin Scorcese", "genre": "Thriller", 
		"quantity": 1, "price": 5999.99
	}`
	assert.JSONEq(t, expectation, w.Body.String())

	mockSvc.AssertExpectations(t)
}

func TestGetTapeByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := NewMockTapeService()
	mockSvc.On("GetTapeByID", "111").Return(nil, false)

	h := handler.NewTapeHandler(mockSvc)

	r := gin.Default()
	r.GET("/tapes/:id", h.GetTapeByID)

	req, _ := http.NewRequest("GET", "/tapes/111", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	expected := `{"error": "tape not found"}`
	assert.JSONEq(t, expected, w.Body.String())

	mockSvc.AssertExpectations(t)
}

func TestCreate_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := NewMockTapeService()

	inputTape := model.Tape{
		ID: "1", Title: "Taxi Driver",
		Director: "Martin Scorsese", Genre: "Thriller",
		Quantity: 1, Price: 5999.99,
	}

	createdTape := &model.Tape{
		ID: "1", Title: "Taxi Driver",
		Director: "Martin Scorsese", Genre: "Thriller",
		Quantity: 1, Price: 5999.99,
	}

	mockSvc.On("Create", inputTape).Return(createdTape, true)

	h := handler.NewTapeHandler(mockSvc)

	r := gin.Default()
	r.POST("/tapes", h.CreateTape)

	data, err := json.Marshal(inputTape)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/tapes", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	jsonExpected := `{
  "id": "1",
  "title": "Taxi Driver",
  "director": "Martin Scorsese",
  "genre": "Thriller",
  "quantity": 1,
  "price": 5999.99
}`
	assert.JSONEq(t, jsonExpected, w.Body.String())

	mockSvc.AssertExpectations(t)
}
