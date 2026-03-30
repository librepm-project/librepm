package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockFilterService struct {
	AllFunc     func() (*[]domain.FilterModel, error)
	ShowFunc    func(id uuid.UUID) (*domain.FilterModel, error)
	CreateFunc  func(filter *domain.FilterModel) error
	UpdateFunc  func(id uuid.UUID, filter *domain.FilterModel) error
	DestroyFunc func(id uuid.UUID) error
}

func (m *MockFilterService) All() (*[]domain.FilterModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.FilterModel{}, nil
}

func (m *MockFilterService) Show(id uuid.UUID) (*domain.FilterModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(id)
	}
	return &domain.FilterModel{}, nil
}

func (m *MockFilterService) Create(filter *domain.FilterModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(filter)
	}
	return nil
}

func (m *MockFilterService) Update(id uuid.UUID, filter *domain.FilterModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, filter)
	}
	return nil
}

func (m *MockFilterService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
