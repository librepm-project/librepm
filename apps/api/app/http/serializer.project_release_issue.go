package http

import (
	"apps/api/app/domain"
	"github.com/google/uuid"
)

type ProjectReleaseIssueResponse struct {
	ID               uuid.UUID              `json:"id"`
	ProjectReleaseID uuid.UUID              `json:"project_release_id"`
	IssueID          uuid.UUID              `json:"issue_id"`
	Issue            IssueResponse          `json:"issue"`
}

type ProjectReleaseIssueSerializer struct{}

func (s ProjectReleaseIssueSerializer) ModelToResponse(pri domain.ProjectReleaseIssueModel) ProjectReleaseIssueResponse {
	return ProjectReleaseIssueResponse{
		ID:               pri.ID,
		ProjectReleaseID: pri.ProjectReleaseID,
		IssueID:          pri.IssueID,
		Issue:            IssueSerializer{}.ModelToResponse(pri.Issue),
	}
}

func (s ProjectReleaseIssueSerializer) ModelToResponseMany(pris []domain.ProjectReleaseIssueModel) []ProjectReleaseIssueResponse {
	responses := make([]ProjectReleaseIssueResponse, len(pris))
	for i, pri := range pris {
		responses[i] = s.ModelToResponse(pri)
	}
	return responses
}
