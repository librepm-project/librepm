package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockProjectTrackerStatusRepository struct {
	AllFunc                     func() (*[]domain.ProjectTrackerStatusModel, error)
	FindByIDFunc                func(id uuid.UUID) (*domain.ProjectTrackerStatusModel, error)
	FindStatusesByProjectIDFunc func(projectID uuid.UUID) (*[]domain.StatusModel, error)
	CreateFunc                  func(pts *domain.ProjectTrackerStatusModel) error
	UpdateFunc                  func(id uuid.UUID, pts *domain.ProjectTrackerStatusModel) error
	DestroyFunc                 func(id uuid.UUID) error
}

func (m *MockProjectTrackerStatusRepository) All() (*[]domain.ProjectTrackerStatusModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ProjectTrackerStatusModel{}, nil
}

func (m *MockProjectTrackerStatusRepository) FindByID(id uuid.UUID) (*domain.ProjectTrackerStatusModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.ProjectTrackerStatusModel{}, nil
}

func (m *MockProjectTrackerStatusRepository) FindStatusesByProjectID(projectID uuid.UUID) (*[]domain.StatusModel, error) {
	if m.FindStatusesByProjectIDFunc != nil {
		return m.FindStatusesByProjectIDFunc(projectID)
	}
	return &[]domain.StatusModel{}, nil
}

func (m *MockProjectTrackerStatusRepository) Create(pts *domain.ProjectTrackerStatusModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(pts)
	}
	return nil
}

func (m *MockProjectTrackerStatusRepository) Update(id uuid.UUID, pts *domain.ProjectTrackerStatusModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, pts)
	}
	return nil
}

func (m *MockProjectTrackerStatusRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
