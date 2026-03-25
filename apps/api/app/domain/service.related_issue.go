package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type RelatedIssueServiceInterface interface {
	Create(issueAID, issueBID uuid.UUID, relationType string) (*RelatedIssueModel, error)
	GetRelatedIssues(issueID uuid.UUID, relationType *string) (*[]RelatedIssueModel, error)
	Delete(relatedIssueID uuid.UUID) error
}

type RelatedIssueService struct {
	RelatedIssueRepository RelatedIssueRepositoryInterface
	IssueRepository        IssueRepositoryInterface
}

func (s RelatedIssueService) Create(issueAID, issueBID uuid.UUID, relationType string) (*RelatedIssueModel, error) {
	// Prevent self-referencing
	if issueAID == issueBID {
		return nil, fmt.Errorf("cannot create self-referencing relationship")
	}

	// Validate both issues exist
	issueA, err := s.IssueRepository.FindByID(issueAID)
	if err != nil {
		return nil, fmt.Errorf("issue A not found: %w", err)
	}

	issueB, err := s.IssueRepository.FindByID(issueBID)
	if err != nil {
		return nil, fmt.Errorf("issue B not found: %w", err)
	}

	// Create relation
	relation := &RelatedIssueModel{
		IssueAID: issueA.ID,
		IssueBID: issueB.ID,
		Type:     relationType,
	}

	if err := s.RelatedIssueRepository.Create(relation); err != nil {
		return nil, fmt.Errorf("failed to create relationship: %w", err)
	}

	return relation, nil
}

func (s RelatedIssueService) GetRelatedIssues(issueID uuid.UUID, relationType *string) (*[]RelatedIssueModel, error) {
	return s.RelatedIssueRepository.FindByIssue(issueID, relationType)
}

func (s RelatedIssueService) Delete(relatedIssueID uuid.UUID) error {
	return s.RelatedIssueRepository.Delete(relatedIssueID)
}
