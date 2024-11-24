package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Board struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type BoardSerializer struct{}

func (s BoardSerializer) SerializeBoard(board domain.BoardModel) Board {
	return Board{
		ID:   board.ID,
		Name: board.Name,
	}
}

func (s BoardSerializer) SerializeBoards(boards []domain.BoardModel) []Board {
	var serialized []Board
	for _, board := range boards {
		serialized = append(serialized, s.SerializeBoard(board))
	}
	return serialized
}
