package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockProjectTrackerRepository struct {
	AllFunc                   func() (*[]domain.ProjectTrackerModel, error)
	FindByIDFunc              func(id uuid.UUID) (*domain.ProjectTrackerModel, error)
	FindTrackersByProjectIDFunc func(projectID uuid.UUID) (*[]domain.TrackerModel, error)
	CreateFunc                func(pt *domain.ProjectTrackerModel) error
	UpdateFunc                func(id uuid.UUID, pt *domain.ProjectTrackerModel) error
	DestroyFunc               func(id uuid.UUID) error
}

func (m *MockProjectTrackerRepository) All() (*[]domain.ProjectTrackerModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ProjectTrackerModel{}, nil
}

func (m *MockProjectTrackerRepository) FindByID(id uuid.UUID) (*domain.ProjectTrackerModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.ProjectTrackerModel{}, nil
}

func (m *MockProjectTrackerRepository) FindTrackersByProjectID(projectID uuid.UUID) (*[]domain.TrackerModel, error) {
	if m.FindTrackersByProjectIDFunc != nil {
		return m.FindTrackersByProjectIDFunc(projectID)
	}
	return &[]domain.TrackerModel{}, nil
}

func (m *MockProjectTrackerRepository) Create(pt *domain.ProjectTrackerModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(pt)
	}
	return nil
}

func (m *MockProjectTrackerRepository) Update(id uuid.UUID, pt *domain.ProjectTrackerModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, pt)
	}
	return nil
}

func (m *MockProjectTrackerRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
