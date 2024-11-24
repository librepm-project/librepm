package domain

import (
	"github.com/google/uuid"
)

type BoardServiceInterface interface {
	All() *[]BoardModel
	Show(board_id uuid.UUID) *BoardModel
	Create(board *BoardModel)
	Update(board_id uuid.UUID, board *BoardModel)
	Destroy(board_id uuid.UUID)
}

type BoardService struct {
	BoardRepository BoardRepositoryInterface
}

func (s BoardService) All() *[]BoardModel {
	boards := s.BoardRepository.All()
	return boards
}

func (s BoardService) Show(board_id uuid.UUID) *BoardModel {
	return s.BoardRepository.FindByID(board_id)
}

func (s BoardService) Create(board *BoardModel) {
	s.BoardRepository.Create(board)
}

func (s BoardService) Update(board_id uuid.UUID, board *BoardModel) {
	s.BoardRepository.Update(board_id, board)
}

func (s BoardService) Destroy(board_id uuid.UUID) {
	s.BoardRepository.Destroy(board_id)
}
