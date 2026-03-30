package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockFilterRepository struct {
	AllFunc      func() (*[]domain.FilterModel, error)
	FindByIDFunc func(id uuid.UUID) (*domain.FilterModel, error)
	FindByNameFunc func(name string) (*domain.FilterModel, error)
	CreateFunc   func(filter *domain.FilterModel) error
	UpdateFunc   func(id uuid.UUID, filter *domain.FilterModel) error
	DestroyFunc  func(id uuid.UUID) error
}

func (m *MockFilterRepository) All() (*[]domain.FilterModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.FilterModel{}, nil
}

func (m *MockFilterRepository) FindByID(id uuid.UUID) (*domain.FilterModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.FilterModel{}, nil
}

func (m *MockFilterRepository) FindByName(name string) (*domain.FilterModel, error) {
	if m.FindByNameFunc != nil {
		return m.FindByNameFunc(name)
	}
	return &domain.FilterModel{}, nil
}

func (m *MockFilterRepository) Create(filter *domain.FilterModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(filter)
	}
	return nil
}

func (m *MockFilterRepository) Update(id uuid.UUID, filter *domain.FilterModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, filter)
	}
	return nil
}

func (m *MockFilterRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
