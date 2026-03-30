package service_test

import (
	"testing"

	"apps/api/app/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockProjectReleaseIssueRepository struct {
	AllFunc                    func() (*[]domain.ProjectReleaseIssueModel, error)
	FindByIDFunc               func(id uuid.UUID) (*domain.ProjectReleaseIssueModel, error)
	FindByProjectReleaseIDFunc func(projectReleaseID uuid.UUID) (*[]domain.ProjectReleaseIssueModel, error)
	CreateFunc                 func(pri *domain.ProjectReleaseIssueModel) error
	DestroyFunc                func(id uuid.UUID) error
}

func (m *MockProjectReleaseIssueRepository) All() (*[]domain.ProjectReleaseIssueModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return nil, nil
}

func (m *MockProjectReleaseIssueRepository) FindByID(id uuid.UUID) (*domain.ProjectReleaseIssueModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}

func (m *MockProjectReleaseIssueRepository) FindByProjectReleaseID(projectReleaseID uuid.UUID) (*[]domain.ProjectReleaseIssueModel, error) {
	if m.FindByProjectReleaseIDFunc != nil {
		return m.FindByProjectReleaseIDFunc(projectReleaseID)
	}
	return nil, nil
}

func (m *MockProjectReleaseIssueRepository) Create(pri *domain.ProjectReleaseIssueModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(pri)
	}
	return nil
}

func (m *MockProjectReleaseIssueRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

func TestProjectReleaseIssueService_All(t *testing.T) {
	priID := uuid.New()
	pris := &[]domain.ProjectReleaseIssueModel{
		{ID: priID, ProjectReleaseID: uuid.New(), IssueID: uuid.New()},
	}

	repo := &MockProjectReleaseIssueRepository{
		AllFunc: func() (*[]domain.ProjectReleaseIssueModel, error) { return pris, nil },
	}
	svc := domain.ProjectReleaseIssueService{ProjectReleaseIssueRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Len(t, *result, 1)
}

func TestProjectReleaseIssueService_Show(t *testing.T) {
	priID := uuid.New()
	pri := &domain.ProjectReleaseIssueModel{ID: priID, ProjectReleaseID: uuid.New(), IssueID: uuid.New()}

	repo := &MockProjectReleaseIssueRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectReleaseIssueModel, error) { return pri, nil },
	}
	svc := domain.ProjectReleaseIssueService{ProjectReleaseIssueRepository: repo}

	result, err := svc.Show(priID)

	assert.NoError(t, err)
	assert.Equal(t, priID, result.ID)
}

func TestProjectReleaseIssueService_FindByProjectReleaseID(t *testing.T) {
	prID := uuid.New()
	pris := &[]domain.ProjectReleaseIssueModel{
		{ID: uuid.New(), ProjectReleaseID: prID, IssueID: uuid.New()},
	}

	repo := &MockProjectReleaseIssueRepository{
		FindByProjectReleaseIDFunc: func(prid uuid.UUID) (*[]domain.ProjectReleaseIssueModel, error) { return pris, nil },
	}
	svc := domain.ProjectReleaseIssueService{ProjectReleaseIssueRepository: repo}

	result, err := svc.FindByProjectReleaseID(prID)

	assert.NoError(t, err)
	assert.Len(t, *result, 1)
}

func TestProjectReleaseIssueService_Create(t *testing.T) {
	pri := &domain.ProjectReleaseIssueModel{ProjectReleaseID: uuid.New(), IssueID: uuid.New()}

	repo := &MockProjectReleaseIssueRepository{
		CreateFunc: func(p *domain.ProjectReleaseIssueModel) error { return nil },
	}
	svc := domain.ProjectReleaseIssueService{ProjectReleaseIssueRepository: repo}

	err := svc.Create(pri)

	assert.NoError(t, err)
}

func TestProjectReleaseIssueService_Destroy(t *testing.T) {
	priID := uuid.New()

	repo := &MockProjectReleaseIssueRepository{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	svc := domain.ProjectReleaseIssueService{ProjectReleaseIssueRepository: repo}

	err := svc.Destroy(priID)

	assert.NoError(t, err)
}
