package repository

import (
	"context"
	"database/sql"
	"sync"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/config"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
)

type TapeRepository interface {
	Save(ctx context.Context, tape *model.Tape) (*model.Tape, error)
	GetAll(ctx context.Context) ([]*model.Tape, error)
	GetByID(ctx context.Context, id int32) (*model.Tape, error)
	GetByPublicID(ctx context.Context, id uuid.UUID) (*model.Tape, error)
	Exists(ctx context.Context, id int32) (bool, error)
	Update(ctx context.Context, updateTape *model.UpdateTape) (*model.Tape, error)
	Delete(ctx context.Context, id int32) error
	DeleteAll(ctx context.Context) error
}

type tapeRepository struct {
	// TODO: mutex is probably not needed
	mu sync.Mutex
	DB *database.Queries
}

func NewTapeRepository() TapeRepository {
	return &tapeRepository{
		DB: config.AppConfig.DB,
	}
}

func (r *tapeRepository) Save(ctx context.Context, tape *model.Tape) (*model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO: Consider doing a function to convert data to and from DAO (similar to the DTO in handler layer)
	tapeParams := database.CreateTapeParams{
		Title:    tape.Title,
		Director: tape.Director,
		Genre:    tape.Genre,
		Quantity: tape.Quantity,
		Price:    tape.Price,
	}

	dbTape, err := r.DB.CreateTape(context.Background(), tapeParams)
	if err != nil {
		return nil, err
	}

	savedTape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID,
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

func (r *tapeRepository) GetAll(ctx context.Context) ([]*model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	dbTapes, err := r.DB.GetTapes(context.Background())
	if err != nil {
		return nil, err
	}
	tapes := make([]*model.Tape, 0)
	for _, tape := range dbTapes {
		t := &model.Tape{
			ID:        tape.ID,
			PublicID:  tape.PublicID,
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

func (r *tapeRepository) GetByID(ctx context.Context, id int32) (*model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbTape, err := r.DB.GetTapeByID(context.Background(), id)
	if err != nil {
		return nil, apperror.ErrTapeNotFound
	}

	tape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID,
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

func (r *tapeRepository) GetByPublicID(ctx context.Context, id uuid.UUID) (*model.Tape, error) {
	dbTape, err := r.DB.GetTapeFromPublicID(ctx, id)
	if err != nil {
		return nil, err
	}
	tape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID,
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

func (r *tapeRepository) Exists(ctx context.Context, id int32) (bool, error) {
	_, err := r.DB.GetTapeByID(ctx, id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *tapeRepository) Update(ctx context.Context, updateTape *model.UpdateTape) (*model.Tape, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	dbUpdateParams := database.UpdateTapeParams{
		ID:       updateTape.ID,
		Title:    toNullString(updateTape.Title),
		Director: toNullString(updateTape.Director),
		Genre:    toNullString(updateTape.Genre),
		Quantity: toNullInt32(updateTape.Quantity),
		Price:    toNullFloat64(updateTape.Price),
	}

	dbTape, err := r.DB.UpdateTape(context.Background(), dbUpdateParams)
	if err != nil {
		return nil, err
	}

	tape := &model.Tape{
		ID:        dbTape.ID,
		PublicID:  dbTape.PublicID,
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

func (r *tapeRepository) Delete(ctx context.Context, id int32) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: Check if tape is in the DB in the service layer
	err := r.DB.DeleteTape(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *tapeRepository) DeleteAll(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.DB.DeleteAllTapes(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Helpers

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func toNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func toNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}
