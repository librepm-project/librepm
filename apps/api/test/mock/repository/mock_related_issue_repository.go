package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockRelatedIssueRepository struct {
	CreateFunc       func(relation *domain.RelatedIssueModel) error
	FindByIssueFunc  func(issueID uuid.UUID, relationType *string) (*[]domain.RelatedIssueModel, error)
	FindByIDFunc     func(id uuid.UUID) (*domain.RelatedIssueModel, error)
	DeleteFunc       func(id uuid.UUID) error
}

func (m *MockRelatedIssueRepository) Create(relation *domain.RelatedIssueModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(relation)
	}
	return nil
}

func (m *MockRelatedIssueRepository) FindByIssue(issueID uuid.UUID, relationType *string) (*[]domain.RelatedIssueModel, error) {
	if m.FindByIssueFunc != nil {
		return m.FindByIssueFunc(issueID, relationType)
	}
	return &[]domain.RelatedIssueModel{}, nil
}

func (m *MockRelatedIssueRepository) FindByID(id uuid.UUID) (*domain.RelatedIssueModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.RelatedIssueModel{}, nil
}

func (m *MockRelatedIssueRepository) Delete(id uuid.UUID) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}
