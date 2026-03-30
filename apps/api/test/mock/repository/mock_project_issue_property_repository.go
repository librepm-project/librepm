package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockProjectIssuePropertyRepository struct {
	FindByProjectIdFunc func(projectID uuid.UUID) (*domain.ProjectIssueProperty, error)
}

func (m *MockProjectIssuePropertyRepository) FindByProjectId(projectID uuid.UUID) (*domain.ProjectIssueProperty, error) {
	if m.FindByProjectIdFunc != nil {
		return m.FindByProjectIdFunc(projectID)
	}
	return &domain.ProjectIssueProperty{}, nil
}
