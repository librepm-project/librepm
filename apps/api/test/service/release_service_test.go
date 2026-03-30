package service_test

import (
	"testing"
	"time"

	"apps/api/app/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockReleaseRepository struct {
	AllFunc    func() (*[]domain.ReleaseModel, error)
	FindByIDFunc func(id uuid.UUID) (*domain.ReleaseModel, error)
	CreateFunc func(release *domain.ReleaseModel) error
	UpdateFunc func(id uuid.UUID, release *domain.ReleaseModel) error
	DestroyFunc func(id uuid.UUID) error
}

func (m *MockReleaseRepository) All() (*[]domain.ReleaseModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return nil, nil
}

func (m *MockReleaseRepository) FindByID(id uuid.UUID) (*domain.ReleaseModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}

func (m *MockReleaseRepository) Create(release *domain.ReleaseModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(release)
	}
	return nil
}

func (m *MockReleaseRepository) Update(id uuid.UUID, release *domain.ReleaseModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, release)
	}
	return nil
}

func (m *MockReleaseRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}

func TestReleaseService_All(t *testing.T) {
	releaseID := uuid.New()
	releases := &[]domain.ReleaseModel{
		{ID: releaseID, Name: "v1.0.0", Status: domain.ReleaseStatusPlanned},
	}

	repo := &MockReleaseRepository{
		AllFunc: func() (*[]domain.ReleaseModel, error) { return releases, nil },
	}
	svc := domain.ReleaseService{ReleaseRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Len(t, *result, 1)
	assert.Equal(t, "v1.0.0", (*result)[0].Name)
}

func TestReleaseService_Show(t *testing.T) {
	releaseID := uuid.New()
	release := &domain.ReleaseModel{ID: releaseID, Name: "v1.0.0", Status: domain.ReleaseStatusPlanned}

	repo := &MockReleaseRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.ReleaseModel, error) { return release, nil },
	}
	svc := domain.ReleaseService{ReleaseRepository: repo}

	result, err := svc.Show(releaseID)

	assert.NoError(t, err)
	assert.Equal(t, "v1.0.0", result.Name)
}

func TestReleaseService_Create(t *testing.T) {
	release := &domain.ReleaseModel{Name: "v1.0.0", Status: domain.ReleaseStatusPlanned}

	repo := &MockReleaseRepository{
		CreateFunc: func(r *domain.ReleaseModel) error { return nil },
	}
	svc := domain.ReleaseService{ReleaseRepository: repo}

	err := svc.Create(release)

	assert.NoError(t, err)
}

func TestReleaseService_Update(t *testing.T) {
	releaseID := uuid.New()
	release := &domain.ReleaseModel{Name: "v1.0.1", Status: domain.ReleaseStatusReleased}

	repo := &MockReleaseRepository{
		UpdateFunc: func(id uuid.UUID, r *domain.ReleaseModel) error { return nil },
	}
	svc := domain.ReleaseService{ReleaseRepository: repo}

	err := svc.Update(releaseID, release)

	assert.NoError(t, err)
}

func TestReleaseService_UpdateWithReleasedAt(t *testing.T) {
	releaseID := uuid.New()
	now := time.Now()
	release := &domain.ReleaseModel{
		Name:       "v1.0.0",
		Status:     domain.ReleaseStatusReleased,
		ReleasedAt: &now,
	}

	repo := &MockReleaseRepository{
		UpdateFunc: func(id uuid.UUID, r *domain.ReleaseModel) error {
			assert.NotNil(t, r.ReleasedAt)
			return nil
		},
	}
	svc := domain.ReleaseService{ReleaseRepository: repo}

	err := svc.Update(releaseID, release)

	assert.NoError(t, err)
}

func TestReleaseService_Destroy(t *testing.T) {
	releaseID := uuid.New()

	repo := &MockReleaseRepository{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	svc := domain.ReleaseService{ReleaseRepository: repo}

	err := svc.Destroy(releaseID)

	assert.NoError(t, err)
}
