package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type IssueRequest struct {
	ProjectID   uuid.UUID `json:"projectId"`
	StatusID    uuid.UUID `json:"statusId"`
	TrackerID   uuid.UUID `json:"trackerId"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
}

type IssueResponse struct {
	ID          uuid.UUID       `json:"id"`
	Key         string          `json:"key"`
	Type        string          `json:"type"`
	Summary     string          `json:"summary"`
	Description string          `json:"description"`
	Project     ProjectResponse `json:"project"`
	Status      StatusResponse  `json:"status"`
	Tracker     TrackerResponse `json:"tracker"`
}

type IssueSerializer struct{}

func (s IssueSerializer) RequestToModel(issue_request IssueRequest) domain.IssueModel {
	return domain.IssueModel{
		ProjectID:   issue_request.ProjectID,
		Key:         issue_request.Key,
		Summary:     issue_request.Summary,
		Description: issue_request.Description,
		StatusID:    issue_request.StatusID,
		TrackerID:   issue_request.TrackerID,
	}
}

func (s IssueSerializer) ModelToResponse(issue domain.IssueModel) IssueResponse {
	return IssueResponse{
		ID:          issue.ID,
		Key:         issue.Key,
		Summary:     issue.Summary,
		Description: issue.Description,
		Project:     ProjectSerializer{}.ModelToResponse(issue.Project),
		Status:      StatusSerializer{}.ModelToResponse(issue.Status),
		Tracker:     TrackerSerializer{}.ModelToResponse(issue.Tracker),
	}
}

func (s IssueSerializer) ModelToResponseMany(issues []domain.IssueModel) []IssueResponse {
	serialized := []IssueResponse{}
	for _, issue := range issues {
		serialized = append(serialized, s.ModelToResponse(issue))
	}
	return serialized
}
