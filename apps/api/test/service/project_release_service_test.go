package service_test

import (
	"testing"

	"apps/api/app/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockProjectReleaseRepository struct {
	AllFunc              func() (*[]domain.ProjectReleaseModel, error)
	FindByIDFunc         func(id uuid.UUID) (*domain.ProjectReleaseModel, error)
	FindByReleaseIDFunc  func(releaseID uuid.UUID) (*[]domain.ProjectReleaseModel, error)
	FindByProjectIDFunc  func(projectID uuid.UUID) (*[]domain.ProjectReleaseModel, error)
	CreateFunc           func(projectRelease *domain.ProjectReleaseModel) error
	UpdateFunc           func(id uuid.UUID, projectRelease *domain.ProjectReleaseModel) error
	DestroyFunc          func(id uuid.UUID) error
}

func (m *MockProjectReleaseRepository) All() (*[]domain.ProjectReleaseModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return nil, nil
}

func (m *MockProjectReleaseRepository) FindByID(id uuid.UUID) (*domain.ProjectReleaseModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}

func (m *MockProjectReleaseRepository) FindByReleaseID(releaseID uuid.UUID) (*[]domain.ProjectReleaseModel, error) {
	if m.FindByReleaseIDFunc != nil {
		return m.FindByReleaseIDFunc(releaseID)
	}
	return nil, nil
}

func (m *MockProjectReleaseRepository) FindByProjectID(projectID uuid.UUID) (*[]domain.ProjectReleaseModel, error) {
	if m.FindByProjectIDFunc != nil {
		return m.FindByProjectIDFunc(projectID)
	}
	return nil, nil
}

func (m *MockProjectReleaseRepository) Create(projectRelease *domain.ProjectReleaseModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(projectRelease)
	}
	return nil
}

func (m *MockProjectReleaseRepository) Update(id uuid.UUID, projectRelease *domain.ProjectReleaseModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, projectRelease)
	}
	return nil
}

func (m *MockProjectReleaseRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

func TestProjectReleaseService_All(t *testing.T) {
	prID := uuid.New()
	prs := &[]domain.ProjectReleaseModel{
		{ID: prID, ReleaseID: uuid.New(), ProjectID: uuid.New()},
	}

	repo := &MockProjectReleaseRepository{
		AllFunc: func() (*[]domain.ProjectReleaseModel, error) { return prs, nil },
	}
	svc := domain.ProjectReleaseService{ProjectReleaseRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Len(t, *result, 1)
}

func TestProjectReleaseService_Show(t *testing.T) {
	prID := uuid.New()
	pr := &domain.ProjectReleaseModel{ID: prID, ReleaseID: uuid.New(), ProjectID: uuid.New()}

	repo := &MockProjectReleaseRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectReleaseModel, error) { return pr, nil },
	}
	svc := domain.ProjectReleaseService{ProjectReleaseRepository: repo}

	result, err := svc.Show(prID)

	assert.NoError(t, err)
	assert.Equal(t, prID, result.ID)
}

func TestProjectReleaseService_FindByReleaseID(t *testing.T) {
	releaseID := uuid.New()
	prs := &[]domain.ProjectReleaseModel{
		{ID: uuid.New(), ReleaseID: releaseID, ProjectID: uuid.New()},
	}

	repo := &MockProjectReleaseRepository{
		FindByReleaseIDFunc: func(rid uuid.UUID) (*[]domain.ProjectReleaseModel, error) { return prs, nil },
	}
	svc := domain.ProjectReleaseService{ProjectReleaseRepository: repo}

	result, err := svc.FindByReleaseID(releaseID)

	assert.NoError(t, err)
	assert.Len(t, *result, 1)
}

func TestProjectReleaseService_FindByProjectID(t *testing.T) {
	projectID := uuid.New()
	prs := &[]domain.ProjectReleaseModel{
		{ID: uuid.New(), ReleaseID: uuid.New(), ProjectID: projectID},
	}

	repo := &MockProjectReleaseRepository{
		FindByProjectIDFunc: func(pid uuid.UUID) (*[]domain.ProjectReleaseModel, error) { return prs, nil },
	}
	svc := domain.ProjectReleaseService{ProjectReleaseRepository: repo}

	result, err := svc.FindByProjectID(projectID)

	assert.NoError(t, err)
	assert.Len(t, *result, 1)
}

func TestProjectReleaseService_Create(t *testing.T) {
	pr := &domain.ProjectReleaseModel{ReleaseID: uuid.New(), ProjectID: uuid.New()}

	repo := &MockProjectReleaseRepository{
		CreateFunc: func(p *domain.ProjectReleaseModel) error { return nil },
	}
	svc := domain.ProjectReleaseService{ProjectReleaseRepository: repo}

	err := svc.Create(pr)

	assert.NoError(t, err)
}

func TestProjectReleaseService_Destroy(t *testing.T) {
	prID := uuid.New()

	repo := &MockProjectReleaseRepository{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	svc := domain.ProjectReleaseService{ProjectReleaseRepository: repo}

	err := svc.Destroy(prID)

	assert.NoError(t, err)
}
