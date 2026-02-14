package service

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type TapeService interface {
	CreateTape(model.Tape) (*model.Tape, error)
	ListTapes() ([]model.Tape, error)
	GetTapeByID(id string) (*model.Tape, error)
	UpdateTape(id string, updated model.UpdatedTape) (*model.Tape, error)
	DeleteTape(id string) error
	DeleteAllTapes() error
}

type tapeService struct {
	repo repository.TapeRepository
}

func NewTapeService(r repository.TapeRepository) TapeService {
	return &tapeService{
		repo: r,
	}
}

// Helper for the Update method
func validateUpdatedTape(id int32, updatedTape model.UpdatedTape) *database.UpdateTapeParams {
	dbUpdatedTapeParams := database.UpdateTapeParams{
		ID: id,
	}

	changes := false

	if updatedTape.Title != nil {
		if *updatedTape.Title == "" {
			return nil
		}
		dbUpdatedTapeParams.Title = sql.NullString{String: *updatedTape.Title, Valid: true}
		changes = true
	} else {
		dbUpdatedTapeParams.Title = sql.NullString{Valid: false}
	}

	if updatedTape.Director != nil {
		if *updatedTape.Director == "" {
			return nil
		}
		dbUpdatedTapeParams.Director = sql.NullString{String: *updatedTape.Director, Valid: true}
		changes = true
	} else {
		dbUpdatedTapeParams.Director = sql.NullString{Valid: false}
	}

	if updatedTape.Genre != nil {
		if *updatedTape.Genre == "" {
			return nil
		}
		dbUpdatedTapeParams.Genre = sql.NullString{String: *updatedTape.Genre, Valid: true}
		changes = true
	} else {
		dbUpdatedTapeParams.Genre = sql.NullString{Valid: false}
	}

	if updatedTape.Quantity != nil {
		// TODO: Decide if I should allow 0 quantity in Update function or not
		// if *updatedTape.Quantity == 0 {
		// 	return nil
		// }
		dbUpdatedTapeParams.Quantity = sql.NullInt32{Int32: *updatedTape.Quantity, Valid: true}
		changes = true
	} else {
		dbUpdatedTapeParams.Quantity = sql.NullInt32{Valid: false}
	}

	if updatedTape.Price != nil {
		if *updatedTape.Price == 0 {
			return nil
		}
		dbUpdatedTapeParams.Price = sql.NullFloat64{Float64: *updatedTape.Price, Valid: true}
		changes = true
	} else {
		dbUpdatedTapeParams.Price = sql.NullFloat64{Valid: false}
	}

	if !changes {
		return nil
	}

	return &dbUpdatedTapeParams
}

// Helper for Create
func validTape(tape model.Tape) bool {
	return tape.Title != "" && tape.Director != "" &&
		tape.Genre != "" && tape.Quantity != 0 && tape.Price != 0
}

// TapeService Methods
//////////////////////

func (s *tapeService) CreateTape(tape model.Tape) (*model.Tape, error) {
	if !validTape(tape) {
		return nil, errors.New("invalid tape fields")
	}

	return s.repo.Save(tape)
}

func (s *tapeService) ListTapes() ([]model.Tape, error) {
	return s.repo.FindAll()
}

func (s *tapeService) GetTapeByID(id string) (*model.Tape, error) {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(int32(tapeID64))
}

func (s *tapeService) UpdateTape(id string, updatedTape model.UpdatedTape) (*model.Tape, error) {
	tapeID64, err := strconv.Atoi(id)
	tapeID32 := int32(tapeID64)
	if err != nil {
		return nil, err
	}

	dbUpdatedParams := validateUpdatedTape(tapeID32, updatedTape)
	if dbUpdatedParams == nil {
		return nil, errors.New("invalid updated tape fields")
	}

	return s.repo.Update(tapeID32, *dbUpdatedParams)
}

func (s *tapeService) DeleteTape(id string) error {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(int32(tapeID64))
}

func (s *tapeService) DeleteAllTapes() error {
	return s.repo.DeleteAll()
}
