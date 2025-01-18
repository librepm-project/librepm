package domain

import (
	"github.com/google/uuid"
)

type BoardServiceInterface interface {
	All() (*[]BoardModel, error)
	Show(board_id uuid.UUID) (*BoardModel, error)
	Create(board *BoardModel) error
	Update(board_id uuid.UUID, board *BoardModel) error
	Destroy(board_id uuid.UUID) error
}

type BoardService struct {
	BoardRepository BoardRepositoryInterface
}

func (s BoardService) All() (*[]BoardModel, error) {
	boards, err := s.BoardRepository.All()
	return boards, err
}

func (s BoardService) Show(board_id uuid.UUID) (*BoardModel, error) {
	return s.BoardRepository.FindByID(board_id)
}

func (s BoardService) Create(board *BoardModel) error {
	return s.BoardRepository.Create(board)
}

func (s BoardService) Update(board_id uuid.UUID, board *BoardModel) error {
	return s.BoardRepository.Update(board_id, board)
}

func (s BoardService) Destroy(board_id uuid.UUID) error {
	return s.BoardRepository.Destroy(board_id)
}
