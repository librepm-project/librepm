package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type BoardRequest struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type BoardResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type BoardSerializer struct{}

func (s BoardSerializer) RequestToModel(board_request BoardRequest) domain.BoardModel {
	return domain.BoardModel{
		Name: board_request.Name,
	}
}

func (s BoardSerializer) ModelToResponse(board domain.BoardModel) BoardResponse {
	return BoardResponse{
		ID:   board.ID,
		Name: board.Name,
	}
}

func (s BoardSerializer) ModelToResponseMany(boards []domain.BoardModel) []BoardResponse {
	serialized := []BoardResponse{}
	for _, board := range boards {
		serialized = append(serialized, s.ModelToResponse(board))
	}
	return serialized
}
