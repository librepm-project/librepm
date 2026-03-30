package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockProjectRepository struct {
	AllFunc                        func() (*[]domain.ProjectModel, error)
	FindByIDFunc                   func(id uuid.UUID) (*domain.ProjectModel, error)
	FindByNameFunc                 func(name string) (*domain.ProjectModel, error)
	CreateFunc                     func(project *domain.ProjectModel) error
	UpdateFunc                     func(id uuid.UUID, project *domain.ProjectModel) error
	DestroyFunc                    func(id uuid.UUID) error
	IncrementLastIssueKeyNumberFunc func(id uuid.UUID) error
}

func (m *MockProjectRepository) All() (*[]domain.ProjectModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ProjectModel{}, nil
}

func (m *MockProjectRepository) FindByID(id uuid.UUID) (*domain.ProjectModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.ProjectModel{}, nil
}

func (m *MockProjectRepository) FindByName(name string) (*domain.ProjectModel, error) {
	if m.FindByNameFunc != nil {
		return m.FindByNameFunc(name)
	}
	return &domain.ProjectModel{}, nil
}

func (m *MockProjectRepository) Create(project *domain.ProjectModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(project)
	}
	return nil
}

func (m *MockProjectRepository) Update(id uuid.UUID, project *domain.ProjectModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, project)
	}
	return nil
}

func (m *MockProjectRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

func (m *MockProjectRepository) IncrementLastIssueKeyNumber(id uuid.UUID) error {
	if m.IncrementLastIssueKeyNumberFunc != nil {
		return m.IncrementLastIssueKeyNumberFunc(id)
	}
	return nil
}
