package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockProjectService struct {
	AllFunc                  func() (*[]domain.ProjectModel, error)
	ShowFunc                 func(id uuid.UUID) (*domain.ProjectModel, error)
	CreateFunc               func(project *domain.ProjectModel) error
	UpdateFunc               func(id uuid.UUID, project *domain.ProjectModel) error
	DestroyFunc              func(id uuid.UUID) error
	ShowIssuePropertyByIdFunc func(id uuid.UUID) (*domain.ProjectIssueProperty, error)
	TrackersByProjectFunc    func(id uuid.UUID) (*[]domain.TrackerModel, error)
	StatusesByProjectFunc    func(id uuid.UUID) (*[]domain.StatusModel, error)
}

func (m *MockProjectService) All() (*[]domain.ProjectModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ProjectModel{}, nil
}

func (m *MockProjectService) Show(id uuid.UUID) (*domain.ProjectModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(id)
	}
	return &domain.ProjectModel{}, nil
}

func (m *MockProjectService) Create(project *domain.ProjectModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(project)
	}
	return nil
}

func (m *MockProjectService) Update(id uuid.UUID, project *domain.ProjectModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, project)
	}
	return nil
}

func (m *MockProjectService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

func (m *MockProjectService) ShowIssuePropertyById(id uuid.UUID) (*domain.ProjectIssueProperty, error) {
	if m.ShowIssuePropertyByIdFunc != nil {
		return m.ShowIssuePropertyByIdFunc(id)
	}
	return &domain.ProjectIssueProperty{}, nil
}

func (m *MockProjectService) TrackersByProject(id uuid.UUID) (*[]domain.TrackerModel, error) {
	if m.TrackersByProjectFunc != nil {
		return m.TrackersByProjectFunc(id)
	}
	return &[]domain.TrackerModel{}, nil
}

func (m *MockProjectService) StatusesByProject(id uuid.UUID) (*[]domain.StatusModel, error) {
	if m.StatusesByProjectFunc != nil {
		return m.StatusesByProjectFunc(id)
	}
	return &[]domain.StatusModel{}, nil
}
