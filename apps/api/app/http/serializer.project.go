package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type ProjectRequest struct {
	Name     string `json:"name"`
	CodeName string `json:"codeName"`
}

type ProjectResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CodeName string    `json:"codeName"`
}

type ProjectSerializer struct{}

func (s ProjectSerializer) RequestToModel(project_request ProjectRequest) domain.ProjectModel {
	return domain.ProjectModel{
		Name:     project_request.Name,
		CodeName: project_request.CodeName,
	}
}

func (s ProjectSerializer) ModelToResponse(project domain.ProjectModel) ProjectResponse {
	return ProjectResponse{
		ID:       project.ID,
		Name:     project.Name,
		CodeName: project.CodeName,
	}
}

func (s ProjectSerializer) ModelToResponseMany(projects []domain.ProjectModel) []ProjectResponse {
	var serialized []ProjectResponse
	for _, project := range projects {
		serialized = append(serialized, s.ModelToResponse(project))
	}
	return serialized
}
