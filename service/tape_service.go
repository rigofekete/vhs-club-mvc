package service

import (
	"database/sql"
	"log"
	"strconv"

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
	DeleteAll() bool
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

func (s *tapeService) Create(tape model.Tape) *model.Tape {
	if !validTape(tape) {
		return nil
	}

	return s.repo.Save(tape)
}

func (s *tapeService) List() []model.Tape {
	return s.repo.FindAll()
}

func (s *tapeService) GetTapeByID(id string) (*model.Tape, bool) {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error parsing id string to uuid: %v", err)
		return nil, false
	}
	return s.repo.FindByID(int32(tapeID64))
}

func (s *tapeService) Update(id string, updatedTape model.UpdatedTape) (*model.Tape, bool) {
	tapeID64, err := strconv.Atoi(id)
	tapeID32 := int32(tapeID64)
	if err != nil {
		log.Printf("error parsing id string to uuid: %v", err)
		return nil, false
	}

	dbUpdatedParams := validateUpdatedTape(tapeID32, updatedTape)
	if dbUpdatedParams == nil {
		log.Print("invalid updated tape")
		return nil, false
	}

	return s.repo.Update(tapeID32, *dbUpdatedParams)
}

func (s *tapeService) Delete(id string) bool {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error parsing id string to uuid: %v", err)
		return false
	}
	return s.repo.Delete(int32(tapeID64))
}

func (s *tapeService) DeleteAll() bool {
	return s.repo.DeleteAllTapes()
}
