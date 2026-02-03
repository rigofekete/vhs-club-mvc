package service_test

import (
	"github.com/rigofekete/vhs-club-mvc/model"
	"github.com/stretchr/testify/mock"
)

type mockTapeRespository struct {
	mock.Mock
}

func (m *mockTapeRespository) Save(tape model.Tape) *model.Tape {
	args := m.Called(tape)
	if t := args.Get(0); t != nil {
		return t.(*model.Tape)
	}
	return nil
}

func (m *mockTapeRespository) FindAll() []model.Tape {
	args := m.Called()
	if tapes := args.Get(0); tapes != nil {
		return tapes.([]model.Tape)
	}
	return nil
}

func (m *mockTapeRespository) FindByID(id string) (*model.Tape, bool) {
	args := m.Called(id)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, args.Bool(1)
}

func (m *mockTapeRespository) Update(id string, updated model.Tape) (*model.Tape, bool) {
	args := m.Called(id, updated)
	if tape := args.Get(0); tape != nil {
		return tape.(*model.Tape), args.Bool(1)
	}
	return nil, false
}

func (m *mockTapeRespository) Delete(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}
