package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockStatusService struct {
	AllFunc     func() (*[]domain.StatusModel, error)
	ShowFunc    func(statusID uuid.UUID) (*domain.StatusModel, error)
	CreateFunc  func(status *domain.StatusModel) error
	UpdateFunc  func(statusID uuid.UUID, status *domain.StatusModel) error
	DestroyFunc func(statusID uuid.UUID) error
}

func (m *MockStatusService) All() (*[]domain.StatusModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.StatusModel{}, nil
}

func (m *MockStatusService) Show(statusID uuid.UUID) (*domain.StatusModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(statusID)
	}
	return &domain.StatusModel{}, nil
}

func (m *MockStatusService) Create(status *domain.StatusModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(status)
	}
	return nil
}

func (m *MockStatusService) Update(statusID uuid.UUID, status *domain.StatusModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(statusID, status)
	}
	return nil
}

func (m *MockStatusService) Destroy(statusID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(statusID)
	}
	return nil
}
