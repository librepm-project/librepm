package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockBoardRepository struct {
	AllFunc      func() (*[]domain.BoardModel, error)
	FindByIDFunc func(id uuid.UUID) (*domain.BoardModel, error)
	CreateFunc   func(board *domain.BoardModel) error
	UpdateFunc   func(id uuid.UUID, board *domain.BoardModel) error
	DestroyFunc  func(id uuid.UUID) error
}

func (m *MockBoardRepository) All() (*[]domain.BoardModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.BoardModel{}, nil
}

func (m *MockBoardRepository) FindByID(id uuid.UUID) (*domain.BoardModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.BoardModel{}, nil
}

func (m *MockBoardRepository) Create(board *domain.BoardModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(board)
	}
	return nil
}

func (m *MockBoardRepository) Update(id uuid.UUID, board *domain.BoardModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, board)
	}
	return nil
}

func (m *MockBoardRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
