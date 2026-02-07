package service

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/database"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type TapeService interface {
	Create(model.Tape) *model.Tape
	List() []model.Tape
	GetTapeByID(id string) (*model.Tape, bool)
	Update(id string, updated model.UpdatedTape) (*model.Tape, bool)
	Delete(id string) bool
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
func validateUpdatedTape(id uuid.UUID, updated model.UpdatedTape) *database.UpdateTapeParams {
	dbUpdatedParams := database.UpdateTapeParams{
		ID: id,
	}

	changes := false

	if updated.Title != nil {
		if *updated.Title == "" {
			return nil
		}
		dbUpdatedParams.Title = sql.NullString{String: *updated.Title, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Title = sql.NullString{Valid: false}
	}

	if updated.Director != nil {
		if *updated.Director == "" {
			return nil
		}
		dbUpdatedParams.Director = sql.NullString{String: *updated.Director, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Director = sql.NullString{Valid: false}
	}

	if updated.Genre != nil {
		if *updated.Genre == "" {
			return nil
		}
		dbUpdatedParams.Genre = sql.NullString{String: *updated.Genre, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Genre = sql.NullString{Valid: false}
	}

	if updated.Quantity != nil {
		// TODO: Decide if I should allow 0 quantity in Update function or not
		// if *updated.Quantity == 0 {
		// 	return nil
		// }
		dbUpdatedParams.Quantity = sql.NullInt32{Int32: *updated.Quantity, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Quantity = sql.NullInt32{Valid: false}
	}

	if updated.Price != nil {
		if *updated.Price == 0 {
			return nil
		}
		dbUpdatedParams.Price = sql.NullFloat64{Float64: *updated.Price, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Price = sql.NullFloat64{Valid: false}
	}

	if !changes {
		return nil
	}

	return &dbUpdatedParams
}

// TapeService Methods
//////////////////////

// TODO: Validate fields to create new tape (e.g. no empty strings, and no 0.00 price)
// NOTE: Consider using the same validateUpdatedTape helper function and the database.UpdateTapeParams type in this function
func (s *tapeService) Create(tape model.Tape) *model.Tape {
	// if !validTape(tape) {
	// 	return nil
	// }
	return s.repo.Save(tape)
}

func (s *tapeService) List() []model.Tape {
	return s.repo.FindAll()
}

func (s *tapeService) GetTapeByID(id string) (*model.Tape, bool) {
	tapeID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error parsing id string to uuid: %v", err)
		return nil, false
	}
	return s.repo.FindByID(tapeID)
}

func (s *tapeService) Update(id string, updated model.UpdatedTape) (*model.Tape, bool) {
	tapeID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error parsing id string to uuid: %v", err)
		return nil, false
	}

	dbUpdatedParams := validateUpdatedTape(tapeID, updated)
	if dbUpdatedParams == nil {
		log.Print("invalid updated tape")
		return nil, false
	}

	return s.repo.Update(tapeID, *dbUpdatedParams)
}

func (s *tapeService) Delete(id string) bool {
	tapeID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error parsing id string to uuid: %v", err)
		return false
	}
	return s.repo.Delete(tapeID)
}
