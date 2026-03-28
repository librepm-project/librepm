package http

import (
	"github.com/google/uuid"
	"apps/api/app/domain"
)

type ProjectRequest struct {
	Name             string `json:"name"`
	CodeName         string `json:"codeName"`
	DefaultStatusID  string `json:"defaultStatusId"`
	DefaultTrackerID string `json:"defaultTrackerId"`
}

type ProjectResponse struct {
	ID               uuid.UUID  `json:"id"`
	Name             string     `json:"name"`
	CodeName         string     `json:"codeName"`
	DefaultStatusID  *uuid.UUID `json:"defaultStatusId"`
	DefaultTrackerID *uuid.UUID `json:"defaultTrackerId"`
}

type ProjectSerializer struct{}

func (s ProjectSerializer) RequestToModel(project_request ProjectRequest) domain.ProjectModel {
	var statusID *uuid.UUID
	if project_request.DefaultStatusID != "" {
		id, err := uuid.Parse(project_request.DefaultStatusID)
		if err == nil {
			statusID = &id
		}
	}

	var trackerID *uuid.UUID
	if project_request.DefaultTrackerID != "" {
		id, err := uuid.Parse(project_request.DefaultTrackerID)
		if err == nil {
			trackerID = &id
		}
	}

	return domain.ProjectModel{
		Name:             project_request.Name,
		CodeName:         project_request.CodeName,
		DefaultStatusID:  statusID,
		DefaultTrackerID: trackerID,
	}
}

func (s ProjectSerializer) ModelToResponse(project domain.ProjectModel) ProjectResponse {
	return ProjectResponse{
		ID:               project.ID,
		Name:             project.Name,
		CodeName:         project.CodeName,
		DefaultStatusID:  project.DefaultStatusID,
		DefaultTrackerID: project.DefaultTrackerID,
	}
}

func (s ProjectSerializer) ModelToResponseMany(projects []domain.ProjectModel) []ProjectResponse {
	serialized := []ProjectResponse{}
	for _, project := range projects {
		serialized = append(serialized, s.ModelToResponse(project))
	}
	return serialized
}
