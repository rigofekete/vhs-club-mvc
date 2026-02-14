package repository

import (
	"context"
	"sync"

	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type TapeRepository interface {
	Save(tape model.Tape) (*model.Tape, error)
	FindAll() ([]model.Tape, error)
	FindByID(id int32) (*model.Tape, error)
	Update(id int32, updated database.UpdateTapeParams) (*model.Tape, error)
	Delete(id int32) error
	DeleteAll() error
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

func (r *tapeRepository) Save(tape model.Tape) (*model.Tape, error) {
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
	// TODO: returning raw sql errors down the chain?
	if err != nil {
		return nil, err
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
	return savedTape, nil
}

func (r *tapeRepository) FindAll() ([]model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	dbTapes, err := r.DB.GetTapes(context.Background())
	if err != nil {
		return nil, err
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
	return tapes, nil
}

func (r *tapeRepository) FindByID(id int32) (*model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.GetTape(context.Background(), id)
	if err != nil {
		return nil, err
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

	return tape, nil
}

func (r *tapeRepository) Update(id int32, updatedTape database.UpdateTapeParams) (*model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.UpdateTape(context.Background(), updatedTape)
	if err != nil {
		return nil, err
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

	return tape, err
}

func (r *tapeRepository) Delete(id int32) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteTape(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *tapeRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteAllTapes(context.Background())
	if err != nil {
		return err
	}
	return nil
}
