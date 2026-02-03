package repository

import (
	"sync"

	"github.com/rigofekete/vhs-club-mvc/model"
)

type TapeRepository interface {
	Save(tape model.Tape) *model.Tape
	FindAll() []model.Tape
	FindByID(id string) (*model.Tape, bool)
	Update(id string, updated model.Tape) (*model.Tape, bool)
	Delete(id string) bool
}

type tapeRepository struct {
	mu    sync.Mutex
	tapes []model.Tape
}

func NewTapeRepository() TapeRepository {
	return &tapeRepository{
		tapes: make([]model.Tape, 0),
	}
}

func (r *tapeRepository) Save(tape model.Tape) *model.Tape {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tapes = append(r.tapes, tape)
	return &tape
}

func (r *tapeRepository) FindAll() []model.Tape {
	r.mu.Lock()
	defer r.mu.Unlock()

	return append([]model.Tape(nil), r.tapes...)
}

func (r *tapeRepository) FindByID(id string) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, tape := range r.tapes {
		if tape.ID == id {
			return &tape, true
		}
	}
	return nil, false
}

func (r *tapeRepository) Update(id string, updated model.Tape) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

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

func (r *tapeRepository) Delete(id string) bool {
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
