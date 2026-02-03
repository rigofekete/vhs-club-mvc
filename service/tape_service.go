package service

import (
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type TapeService interface {
	Create(model.Tape) *model.Tape
	List() []model.Tape
	GetTapeByID(id string) (*model.Tape, bool)
	Update(id string, updated model.Tape) (*model.Tape, bool)
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

func (s *tapeService) Create(tape model.Tape) *model.Tape {
	if !validTape(tape) {
		return nil
	}
	return s.repo.Save(tape)
}

// Helper for Create
func validTape(tape model.Tape) bool {
	if tape.Title == "" || tape.Director == "" || tape.Genre == "" || tape.Quantity == 0 || tape.Price == 0 {
		return false
	}
	return true
}

func (s *tapeService) List() []model.Tape {
	return s.repo.FindAll()
}

func (s *tapeService) GetTapeByID(id string) (*model.Tape, bool) {
	return s.repo.FindByID(id)
}

func (s *tapeService) Update(id string, updated model.Tape) (*model.Tape, bool) {
	return s.repo.Update(id, updated)
}

func (s *tapeService) Delete(id string) bool {
	return s.repo.Delete(id)
}
