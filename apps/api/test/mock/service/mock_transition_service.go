package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockTransitionService struct {
	AllFunc     func() (*[]domain.TransitionModel, error)
	ShowFunc    func(transitionID uuid.UUID) (*domain.TransitionModel, error)
	CreateFunc  func(transition *domain.TransitionModel) error
	UpdateFunc  func(transitionID uuid.UUID, transition *domain.TransitionModel) error
	DestroyFunc func(transitionID uuid.UUID) error
}

func (m *MockTransitionService) All() (*[]domain.TransitionModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.TransitionModel{}, nil
}

func (m *MockTransitionService) Show(transitionID uuid.UUID) (*domain.TransitionModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(transitionID)
	}
	return &domain.TransitionModel{}, nil
}

func (m *MockTransitionService) Create(transition *domain.TransitionModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(transition)
	}
	return nil
}

func (m *MockTransitionService) Update(transitionID uuid.UUID, transition *domain.TransitionModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(transitionID, transition)
	}
	return nil
}

func (m *MockTransitionService) Destroy(transitionID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(transitionID)
	}
	return nil
}
