package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockStatusRepository struct {
	AllFunc        func() (*[]domain.StatusModel, error)
	FindByIDFunc   func(id uuid.UUID) (*domain.StatusModel, error)
	FindByNameFunc func(name string) (*domain.StatusModel, error)
	CreateFunc     func(status *domain.StatusModel) error
	UpdateFunc     func(id uuid.UUID, status *domain.StatusModel) error
	DestroyFunc    func(id uuid.UUID) error
}

func (m *MockStatusRepository) All() (*[]domain.StatusModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.StatusModel{}, nil
}

func (m *MockStatusRepository) FindByID(id uuid.UUID) (*domain.StatusModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.StatusModel{}, nil
}

func (m *MockStatusRepository) FindByName(name string) (*domain.StatusModel, error) {
	if m.FindByNameFunc != nil {
		return m.FindByNameFunc(name)
	}
	return &domain.StatusModel{}, nil
}

func (m *MockStatusRepository) Create(status *domain.StatusModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(status)
	}
	return nil
}

func (m *MockStatusRepository) Update(id uuid.UUID, status *domain.StatusModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, status)
	}
	return nil
}

func (m *MockStatusRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
