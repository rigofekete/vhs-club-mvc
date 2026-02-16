package handler

import "github.com/rigofekete/vhs-club-mvc/model"

func (r *CreateTapeRequest) ToModel() *model.Tape {
	return &model.Tape{
		Title:    r.Title,
		Director: r.Director,
		Genre:    r.Genre,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}

func TapeSingleResponse(tape *model.Tape) *TapeResponse {
	return &TapeResponse{
		PublicID:  tape.PublicID,
		CreatedAt: tape.CreatedAt,
		UpdatedAt: tape.UpdatedAt,
		Title:     tape.Title,
		Director:  tape.Director,
		Genre:     tape.Genre,
		Quantity:  tape.Quantity,
		Price:     tape.Price,
	}
}

func TapeListResponse(tapes []*model.Tape) []*TapeResponse {
	tapeList := make([]*TapeResponse, len(tapes))
	for i, tape := range tapes {
		tapeList[i] = TapeSingleResponse(tape)
	}
	return tapeList
}

func (r UpdateTapeRequest) ToModel() *model.UpdateTape {
	return &model.UpdateTape{
		Title:    r.Title,
		Director: r.Director,
		Genre:    r.Genre,
		Quantity: r.Quantity,
		Price:    r.Price,
	}
}

func TapeUpdateResponse(tape *model.Tape) *TapeResponse {
	return &TapeResponse{
		PublicID:  tape.PublicID,
		CreatedAt: tape.CreatedAt,
		UpdatedAt: tape.UpdatedAt,
		Title:     tape.Title,
		Director:  tape.Director,
		Genre:     tape.Genre,
		Quantity:  tape.Quantity,
		Price:     tape.Price,
	}
}
