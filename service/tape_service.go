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
		dbUpdatedParams.Title = sql.NullString{String: *updated.Title, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Title = sql.NullString{Valid: false}
	}

	if updated.Director != nil {
		dbUpdatedParams.Director = sql.NullString{String: *updated.Director, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Director = sql.NullString{Valid: false}
	}

	if updated.Genre != nil {
		dbUpdatedParams.Genre = sql.NullString{String: *updated.Genre, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Genre = sql.NullString{Valid: false}
	}

	if updated.Quantity != nil {
		dbUpdatedParams.Quantity = sql.NullInt32{Int32: *updated.Quantity, Valid: true}
		changes = true
	} else {
		dbUpdatedParams.Quantity = sql.NullInt32{Valid: false}
	}

	if updated.Price != nil {
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

func (s *tapeService) Create(tape model.Tape) *model.Tape {
	// if !validTape(tape) {
	// 	return nil
	// }
	return s.repo.Save(tape)
}

// NOTE: With the DB we won't need this any longer
// Helper for Create
// func validTape(tape model.Tape) bool {
// 	if tape.Title == "" || tape.Director == "" || tape.Genre == "" || tape.Quantity == 0 || tape.Price == 0 {
// 		return false
// 	}
// 	return true
// }

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
