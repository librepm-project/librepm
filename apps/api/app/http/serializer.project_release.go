package http

import (
	"apps/api/app/domain"
	"github.com/google/uuid"
)

type ProjectReleaseResponse struct {
	ID        uuid.UUID        `json:"id"`
	ReleaseID uuid.UUID        `json:"release_id"`
	ProjectID uuid.UUID        `json:"project_id"`
	Release   ReleaseResponse  `json:"release"`
	Project   ProjectResponse  `json:"project"`
}

type ProjectReleaseSerializer struct{}

func (s ProjectReleaseSerializer) ModelToResponse(pr domain.ProjectReleaseModel) ProjectReleaseResponse {
	return ProjectReleaseResponse{
		ID:        pr.ID,
		ReleaseID: pr.ReleaseID,
		ProjectID: pr.ProjectID,
		Release:   ReleaseSerializer{}.ModelToResponse(pr.Release),
		Project:   ProjectSerializer{}.ModelToResponse(pr.Project),
	}
}

func (s ProjectReleaseSerializer) ModelToResponseMany(prs []domain.ProjectReleaseModel) []ProjectReleaseResponse {
	responses := make([]ProjectReleaseResponse, len(prs))
	for i, pr := range prs {
		responses[i] = s.ModelToResponse(pr)
	}
	return responses
}
