package repository

import (
	"context"
	"log"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type TapeRepository interface {
	Save(tape model.Tape) *model.Tape
	FindAll() []model.Tape
	FindByID(id int32) (*model.Tape, bool)
	Update(id int32, updated database.UpdateTapeParams) (*model.Tape, bool)
	Delete(id int32) bool
	DeleteAllTapes() bool
}

type tapeRepository struct {
	mu sync.Mutex
	DB *database.Queries
}

func NewTapeRepository() TapeRepository {
	return &tapeRepository{
		DB: config.AppConfig.DB,
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
		// TODO: Should we return the err together with the object pointer?
		return nil
	}

	savedTape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID.UUID,
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
		// TODO: Should we return the err together with the object pointer?
		return nil
	}
	tapes := make([]model.Tape, 0)
	for _, tape := range dbTapes {
		t := model.Tape{
			ID:        tape.ID,
			PublicID:  tape.PublicID.UUID,
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

func (r *tapeRepository) FindByID(id int32) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.GetTape(context.Background(), id)
	if err != nil {
		log.Printf("error from DB.GetTape request: %v", err)
		return nil, false
	}

	tape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID.UUID,
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

func (r *tapeRepository) Update(id int32, updatedTape database.UpdateTapeParams) (*model.Tape, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.UpdateTape(context.Background(), updatedTape)
	if err != nil {
		log.Printf("error updating tape in the db: %v", err)
		return nil, false
	}

	tape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID.UUID,
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

func (r *tapeRepository) Delete(id int32) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteTape(context.Background(), id)
	if err != nil {
		log.Printf("error deleting tape from the db: %v", err)
		return false
	}
	return true
}

func (r *tapeRepository) DeleteAllTapes() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteAllTapes(context.Background())
	if err != nil {
		log.Printf("error deleting all tapes from the db: %v", err)
		return false
	}
	return true
}
