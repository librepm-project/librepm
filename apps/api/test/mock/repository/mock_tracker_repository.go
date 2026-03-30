package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockTrackerRepository struct {
	AllFunc        func() (*[]domain.TrackerModel, error)
	FindByIDFunc   func(id uuid.UUID) (*domain.TrackerModel, error)
	FindByNameFunc func(name string) (*domain.TrackerModel, error)
	CreateFunc     func(tracker *domain.TrackerModel) error
	UpdateFunc     func(id uuid.UUID, tracker *domain.TrackerModel) error
	DestroyFunc    func(id uuid.UUID) error
}

func (m *MockTrackerRepository) All() (*[]domain.TrackerModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.TrackerModel{}, nil
}

func (m *MockTrackerRepository) FindByID(id uuid.UUID) (*domain.TrackerModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.TrackerModel{}, nil
}

func (m *MockTrackerRepository) FindByName(name string) (*domain.TrackerModel, error) {
	if m.FindByNameFunc != nil {
		return m.FindByNameFunc(name)
	}
	return &domain.TrackerModel{}, nil
}

func (m *MockTrackerRepository) Create(tracker *domain.TrackerModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(tracker)
	}
	return nil
}

func (m *MockTrackerRepository) Update(id uuid.UUID, tracker *domain.TrackerModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, tracker)
	}
	return nil
}

func (m *MockTrackerRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
