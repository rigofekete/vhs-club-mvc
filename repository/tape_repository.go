package repository

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type TapeRepository interface {
	Save(tape model.Tape) *model.Tape
	FindAll() []model.Tape
	FindByID(id uuid.UUID) (*model.Tape, bool)
	Update(id uuid.UUID, updated database.UpdateTapeParams) (*model.Tape, bool)
	Delete(id uuid.UUID) bool
}

type tapeRepository struct {
	mu    sync.Mutex
	tapes []model.Tape
	DB    *database.Queries
}

func NewTapeRepository() TapeRepository {
	return &tapeRepository{
		DB: config.AppConfig.DB,
		// tapes: make([]model.Tape, 0),
	}
}

func (r *tapeRepository) Save(tape model.Tape) *model.Tape {
	r.mu.Lock()
	defer r.mu.Unlock()
	tapeParams := database.CreateTapeParams{
		Title:    tape.Title,
		Director: tape.Director,
		Genre:    tape.Genre,
		Quantity: tape.Quantity,
		Price:    tape.Price,
	}

	dbTape, err := r.DB.CreateTape(context.Background(), tapeParams)
	if err != nil {
		log.Fatalf("error creating tape in the db: %v", err)
	}
	savedTape := &model.Tape{
		ID:        dbTape.ID,
		CreatedAt: dbTape.CreatedAt,
		UpdatedAt: dbTape.UpdatedAt,
		Title:     dbTape.Title,
		Director:  dbTape.Director,
		Genre:     dbTape.Genre,
		Quantity:  dbTape.Quantity,
		Price:     dbTape.Price,
	}
	return savedTape
}

func (r *tapeRepository) FindAll() []model.Tape {
	r.mu.Lock()
	defer r.mu.Unlock()
	dbTapes, err := r.DB.GetTapes(context.Background())
	if err != nil {
		log.Fatalf("error getting tapes list from the db: %v", err)
	}
	tapes := make([]model.Tape, 0)
	for _, tape := range dbTapes {
		t := model.Tape{
			ID:        tape.ID,
			CreatedAt: tape.CreatedAt,
			UpdatedAt: tape.UpdatedAt,
			Title:     tape.Title,
			Director:  tape.Director,
			Genre:     tape.Genre,
			Quantity:  tape.Quantity,
			Price:     tape.Price,
		}
		tapes = append(tapes, t)
	}
	return tapes
}

func (r *tapeRepository) FindByID(id uuid.UUID) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.GetTape(context.Background(), id)
	if err != nil {
		// Debug error print
		// fmt.Errorf("error from DB.GetTape request: %v", err)
		return nil, false
	}

	tape := &model.Tape{
		ID:        dbTape.ID,
		CreatedAt: dbTape.CreatedAt,
		UpdatedAt: dbTape.UpdatedAt,
		Title:     dbTape.Title,
		Director:  dbTape.Director,
		Genre:     dbTape.Genre,
		Quantity:  dbTape.Quantity,
		Price:     dbTape.Price,
	}

	return tape, true
}

func (r *tapeRepository) Update(id uuid.UUID, updated database.UpdateTapeParams) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.UpdateTape(context.Background(), updated)
	if err != nil {
		log.Printf("error updating tape in the db: %v", err)
		return nil, false
	}

	tape := &model.Tape{
		ID:        dbTape.ID,
		CreatedAt: dbTape.CreatedAt,
		UpdatedAt: dbTape.UpdatedAt,
		Title:     dbTape.Title,
		Director:  dbTape.Director,
		Genre:     dbTape.Genre,
		Quantity:  dbTape.Quantity,
		Price:     dbTape.Price,
	}

	return tape, true
}

func (r *tapeRepository) Delete(id uuid.UUID) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, tape := range r.tapes {
		if tape.ID == id {
			r.tapes = append(r.tapes[:i], r.tapes[i+1:]...)
			return true
		}
	}
	return false
}
