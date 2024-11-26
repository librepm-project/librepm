package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Issue struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"projectId"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

type IssueSerializer struct{}

func (s IssueSerializer) SerializeIssue(issue domain.IssueModel) Issue {
	return Issue{
		ID:        issue.ID,
		ProjectID: issue.ProjectID,
		Key:       issue.Key,

		Summary:     issue.Summary,
		Description: issue.Description,
		Status:      issue.IssueStatus.Name,
	}
}

func (s IssueSerializer) SerializeIssues(issues []domain.IssueModel) []Issue {
	var serialized []Issue
	for _, issue := range issues {
		serialized = append(serialized, s.SerializeIssue(issue))
	}
	return serialized
}
