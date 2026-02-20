package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rigofekete/vhs-club-mvc/internal/apperror"
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/rigofekete/vhs-club-mvc/repository"
)

type TapeService interface {
	CreateTape(ctx context.Context, tape *model.Tape) (*model.Tape, error)
	GetAllTapes(ctx context.Context) ([]*model.Tape, error)
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
	dbTapes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, dbTape := range dbTapes {
		// TODO: Make Title and Director composite UNIQUE in sql
		if dbTape.Title == tape.Title {
			return nil, apperror.ErrTapeExists
		}
	}

	return s.repo.Save(ctx, tape)
}

func (s *tapeService) GetAllTapes(ctx context.Context) ([]*model.Tape, error) {
	return s.repo.GetAll(ctx)
}

func (s *tapeService) GetTapeByID(ctx context.Context, id string) (*model.Tape, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// TODO: Is it good practice to just get the tape directly with the public ID from the repo?
	tape, err := s.repo.GetByPublicID(ctx, idUUID)
	if err != nil {
		return nil, apperror.ErrTapeNotFound
	}

	return tape, nil
}

func (s *tapeService) UpdateTape(ctx context.Context, id string, updateTape *model.UpdateTape) (*model.Tape, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	tape, err := s.repo.GetByPublicID(ctx, idUUID)
	if err != nil {
		return nil, apperror.ErrTapeNotFound
	}

	updateTape.ID = tape.ID

	return s.repo.Update(ctx, updateTape)
}

func (s *tapeService) DeleteTape(ctx context.Context, id string) error {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	tape, err := s.repo.GetByPublicID(ctx, idUUID)
	if err != nil {
		return apperror.ErrTapeNotFound
	}

	return s.repo.Delete(ctx, tape.ID)
}

func (s *tapeService) DeleteAllTapes(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}
