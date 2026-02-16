package service

import (
	"context"
	"strconv"

	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type TapeService interface {
	CreateTape(ctx context.Context, tape *model.Tape) (*model.Tape, error)
	ListTapes(ctx context.Context) ([]*model.Tape, error)
	GetTapeByID(ctx context.Context, id string) (*model.Tape, error)
	UpdateTape(ctx context.Context, id string, updated *model.UpdateTape) (*model.Tape, error)
	DeleteTape(ctx context.Context, id string) error
	DeleteAllTapes(ctx context.Context) error
}

type tapeService struct {
	repo repository.TapeRepository
}

func NewTapeService(r repository.TapeRepository) TapeService {
	return &tapeService{
		repo: r,
	}
}

func (s *tapeService) CreateTape(ctx context.Context, tape *model.Tape) (*model.Tape, error) {
	return s.repo.Save(ctx, tape)
}

func (s *tapeService) ListTapes(ctx context.Context) ([]*model.Tape, error) {
	return s.repo.FindAll(ctx)
}

func (s *tapeService) GetTapeByID(ctx context.Context, id string) (*model.Tape, error) {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, int32(tapeID64))
}

func (s *tapeService) UpdateTape(ctx context.Context, id string, updateTape *model.UpdateTape) (*model.Tape, error) {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	updateTape.ID = int32(tapeID64)

	// TODO: Good practice to do this here?
	exists, err := s.repo.Exists(ctx, updateTape.ID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, apperror.ErrTapeNotFound
	}

	return s.repo.Update(ctx, updateTape)
}

// TODO: Which id to pass to delete tape?
func (s *tapeService) DeleteTape(ctx context.Context, id string) error {
	tapeID64, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, int32(tapeID64))
}

func (s *tapeService) DeleteAllTapes(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
