package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockBoardService struct {
	AllFunc     func() (*[]domain.BoardModel, error)
	ShowFunc    func(boardID uuid.UUID) (*domain.BoardModel, error)
	CreateFunc  func(board *domain.BoardModel) error
	UpdateFunc  func(boardID uuid.UUID, board *domain.BoardModel) error
	DestroyFunc func(boardID uuid.UUID) error
}

func (m *MockBoardService) All() (*[]domain.BoardModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.BoardModel{}, nil
}

func (m *MockBoardService) Show(boardID uuid.UUID) (*domain.BoardModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(boardID)
	}
	return &domain.BoardModel{}, nil
}

func (m *MockBoardService) Create(board *domain.BoardModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(board)
	}
	return nil
}

func (m *MockBoardService) Update(boardID uuid.UUID, board *domain.BoardModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(boardID, board)
	}
	return nil
}

func (m *MockBoardService) Destroy(boardID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(boardID)
	}
	return nil
}
