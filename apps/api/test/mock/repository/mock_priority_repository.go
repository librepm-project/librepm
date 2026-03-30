package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockPriorityRepository struct {
	AllFunc      func() (*[]domain.PriorityModel, error)
	FindByIDFunc func(id uuid.UUID) (*domain.PriorityModel, error)
	CreateFunc   func(priority *domain.PriorityModel) error
	UpdateFunc   func(id uuid.UUID, priority *domain.PriorityModel) error
	DestroyFunc  func(id uuid.UUID) error
}

func (m *MockPriorityRepository) All() (*[]domain.PriorityModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.PriorityModel{}, nil
}

func (m *MockPriorityRepository) FindByID(id uuid.UUID) (*domain.PriorityModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.PriorityModel{}, nil
}

func (m *MockPriorityRepository) Create(priority *domain.PriorityModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(priority)
	}
	return nil
}

func (m *MockPriorityRepository) Update(id uuid.UUID, priority *domain.PriorityModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, priority)
	}
	return nil
}

func (m *MockPriorityRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
