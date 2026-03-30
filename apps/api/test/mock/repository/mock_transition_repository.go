package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockTransitionRepository struct {
	AllFunc      func() (*[]domain.TransitionModel, error)
	FindByIDFunc func(id uuid.UUID) (*domain.TransitionModel, error)
	CreateFunc   func(transition *domain.TransitionModel) error
	UpdateFunc   func(id uuid.UUID, transition *domain.TransitionModel) error
	DestroyFunc  func(id uuid.UUID) error
}

func (m *MockTransitionRepository) All() (*[]domain.TransitionModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.TransitionModel{}, nil
}

func (m *MockTransitionRepository) FindByID(id uuid.UUID) (*domain.TransitionModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.TransitionModel{}, nil
}

func (m *MockTransitionRepository) Create(transition *domain.TransitionModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(transition)
	}
	return nil
}

func (m *MockTransitionRepository) Update(id uuid.UUID, transition *domain.TransitionModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, transition)
	}
	return nil
}

func (m *MockTransitionRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
