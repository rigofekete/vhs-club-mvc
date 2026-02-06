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
	Update(id uuid.UUID, updated model.Tape) (*model.Tape, bool)
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
		Quantity: int32(tape.Quantity),
		Price:    int32(tape.Price),
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
		Quantity:  int(dbTape.Quantity),
		Price:     float64(dbTape.Price),
	}
	// r.tapes = append(r.tapes, tape)
	return savedTape
}

func (r *tapeRepository) FindAll() []model.Tape {
	r.mu.Lock()
	defer r.mu.Unlock()

	return append([]model.Tape(nil), r.tapes...)
}

func (r *tapeRepository) FindByID(id uuid.UUID) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, tape := range r.tapes {
		if tape.ID == id {
			return &tape, true
		}
	}
	return nil, false
}

func (r *tapeRepository) Update(id uuid.UUID, updated model.Tape) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: Shouldn't this data validation be done in the service layer?
	for i, tape := range r.tapes {
		if tape.ID == id {
			if updated.Title != "" {
				r.tapes[i].Title = updated.Title
			}
			if updated.Director != "" {
				r.tapes[i].Director = updated.Director
			}
			if updated.Genre != "" {
				r.tapes[i].Genre = updated.Genre
			}
			if updated.Quantity != 0 {
				r.tapes[i].Quantity = updated.Quantity
			}
			if updated.Price != 0 {
				r.tapes[i].Price = updated.Price
			}
			return &updated, true
		}
	}
	return nil, false
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
