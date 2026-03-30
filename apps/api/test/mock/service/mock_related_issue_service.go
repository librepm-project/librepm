package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockRelatedIssueService struct {
	CreateFunc          func(issueAID, issueBID uuid.UUID, relationType string) (*domain.RelatedIssueModel, error)
	GetRelatedIssuesFunc func(issueID uuid.UUID, relationType *string) (*[]domain.RelatedIssueModel, error)
	DeleteFunc          func(relatedIssueID uuid.UUID) error
	FindByIDFunc        func(relatedIssueID uuid.UUID) (*domain.RelatedIssueModel, error)
}

func (m *MockRelatedIssueService) Create(issueAID, issueBID uuid.UUID, relationType string) (*domain.RelatedIssueModel, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(issueAID, issueBID, relationType)
	}
	return &domain.RelatedIssueModel{}, nil
}

func (m *MockRelatedIssueService) GetRelatedIssues(issueID uuid.UUID, relationType *string) (*[]domain.RelatedIssueModel, error) {
	if m.GetRelatedIssuesFunc != nil {
		return m.GetRelatedIssuesFunc(issueID, relationType)
	}
	return &[]domain.RelatedIssueModel{}, nil
}

func (m *MockRelatedIssueService) Delete(relatedIssueID uuid.UUID) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(relatedIssueID)
	}
	return nil
}

func (m *MockRelatedIssueService) FindByID(relatedIssueID uuid.UUID) (*domain.RelatedIssueModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(relatedIssueID)
	}
	return &domain.RelatedIssueModel{}, nil
}
