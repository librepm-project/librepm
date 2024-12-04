package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type IssueRequest struct {
	ProjectID   uuid.UUID `json:"projectId"`
	StatusID    uuid.UUID `json:"statusId"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
}

type IssueResponse struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"projectId"`
	StatusID    uuid.UUID `json:"statusId"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
}

type IssueSerializer struct{}

func (s IssueSerializer) RequestToModel(issue_request IssueRequest) domain.IssueModel {
	return domain.IssueModel{
		ProjectID:   issue_request.ProjectID,
		Key:         issue_request.Key,
		Summary:     issue_request.Summary,
		Description: issue_request.Description,
		StatusID:    issue_request.StatusID,
	}
}

func (s IssueSerializer) ModelToResponse(issue domain.IssueModel) IssueResponse {
	return IssueResponse{
		ID:          issue.ID,
		ProjectID:   issue.ProjectID,
		Key:         issue.Key,
		Summary:     issue.Summary,
		Description: issue.Description,
		StatusID:    issue.StatusID,
	}
}

func (s IssueSerializer) ModelToResponseMany(issues []domain.IssueModel) []IssueResponse {
	var serialized []IssueResponse
	for _, issue := range issues {
		serialized = append(serialized, s.ModelToResponse(issue))
	}
	return serialized
}
