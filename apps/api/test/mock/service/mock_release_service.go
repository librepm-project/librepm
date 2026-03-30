package mocksvc

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type MockReleaseService struct {
	AllFunc    func() (*[]domain.ReleaseModel, error)
	ShowFunc   func(id uuid.UUID) (*domain.ReleaseModel, error)
	CreateFunc func(release *domain.ReleaseModel) error
	UpdateFunc func(id uuid.UUID, release *domain.ReleaseModel) error
	DestroyFunc func(id uuid.UUID) error
}

func (m *MockReleaseService) All() (*[]domain.ReleaseModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ReleaseModel{}, nil
}

func (m *MockReleaseService) Show(id uuid.UUID) (*domain.ReleaseModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(id)
	}
	return &domain.ReleaseModel{}, nil
}

func (m *MockReleaseService) Create(release *domain.ReleaseModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(release)
	}
	return nil
}

func (m *MockReleaseService) Update(id uuid.UUID, release *domain.ReleaseModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, release)
	}
	return nil
}

func (m *MockReleaseService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

type MockProjectReleaseService struct {
	AllFunc              func() (*[]domain.ProjectReleaseModel, error)
	ShowFunc             func(id uuid.UUID) (*domain.ProjectReleaseModel, error)
	FindByReleaseIDFunc  func(releaseID uuid.UUID) (*[]domain.ProjectReleaseModel, error)
	FindByProjectIDFunc  func(projectID uuid.UUID) (*[]domain.ProjectReleaseModel, error)
	CreateFunc           func(projectRelease *domain.ProjectReleaseModel) error
	UpdateFunc           func(id uuid.UUID, projectRelease *domain.ProjectReleaseModel) error
	DestroyFunc          func(id uuid.UUID) error
}

func (m *MockProjectReleaseService) All() (*[]domain.ProjectReleaseModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ProjectReleaseModel{}, nil
}

func (m *MockProjectReleaseService) Show(id uuid.UUID) (*domain.ProjectReleaseModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(id)
	}
	return &domain.ProjectReleaseModel{}, nil
}

func (m *MockProjectReleaseService) FindByReleaseID(releaseID uuid.UUID) (*[]domain.ProjectReleaseModel, error) {
	if m.FindByReleaseIDFunc != nil {
		return m.FindByReleaseIDFunc(releaseID)
	}
	return &[]domain.ProjectReleaseModel{}, nil
}

func (m *MockProjectReleaseService) FindByProjectID(projectID uuid.UUID) (*[]domain.ProjectReleaseModel, error) {
	if m.FindByProjectIDFunc != nil {
		return m.FindByProjectIDFunc(projectID)
	}
	return &[]domain.ProjectReleaseModel{}, nil
}

func (m *MockProjectReleaseService) Create(projectRelease *domain.ProjectReleaseModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(projectRelease)
	}
	return nil
}

func (m *MockProjectReleaseService) Update(id uuid.UUID, projectRelease *domain.ProjectReleaseModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, projectRelease)
	}
	return nil
}

func (m *MockProjectReleaseService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

type MockProjectReleaseIssueService struct {
	AllFunc                         func() (*[]domain.ProjectReleaseIssueModel, error)
	ShowFunc                        func(id uuid.UUID) (*domain.ProjectReleaseIssueModel, error)
	FindByProjectReleaseIDFunc      func(projectReleaseID uuid.UUID) (*[]domain.ProjectReleaseIssueModel, error)
	FindByIssueIDFunc               func(issueID uuid.UUID) (*domain.ProjectReleaseIssueModel, error)
	CreateFunc                      func(pri *domain.ProjectReleaseIssueModel) error
	DestroyFunc                     func(id uuid.UUID) error
}

func (m *MockProjectReleaseIssueService) All() (*[]domain.ProjectReleaseIssueModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.ProjectReleaseIssueModel{}, nil
}

func (m *MockProjectReleaseIssueService) Show(id uuid.UUID) (*domain.ProjectReleaseIssueModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(id)
	}
	return &domain.ProjectReleaseIssueModel{}, nil
}

func (m *MockProjectReleaseIssueService) FindByProjectReleaseID(projectReleaseID uuid.UUID) (*[]domain.ProjectReleaseIssueModel, error) {
	if m.FindByProjectReleaseIDFunc != nil {
		return m.FindByProjectReleaseIDFunc(projectReleaseID)
	}
	return &[]domain.ProjectReleaseIssueModel{}, nil
}

func (m *MockProjectReleaseIssueService) FindByIssueID(issueID uuid.UUID) (*domain.ProjectReleaseIssueModel, error) {
	if m.FindByIssueIDFunc != nil {
		return m.FindByIssueIDFunc(issueID)
	}
	return &domain.ProjectReleaseIssueModel{}, nil
}

func (m *MockProjectReleaseIssueService) Create(pri *domain.ProjectReleaseIssueModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(pri)
	}
	return nil
}

func (m *MockProjectReleaseIssueService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
