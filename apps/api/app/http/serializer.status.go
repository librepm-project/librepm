package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type StatusRequest struct {
	Name string `json:"name"`
}

type StatusResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type StatusSerializer struct{}

func (s StatusSerializer) RequestToModel(status_request StatusRequest) domain.StatusModel {
	return domain.StatusModel{
		Name: status_request.Name,
	}
}

func (s StatusSerializer) ModelToResponse(status domain.StatusModel) StatusResponse {
	return StatusResponse{
		ID:   status.ID,
		Name: status.Name,
	}
}

func (s StatusSerializer) ModelToResponseMany(statuses []domain.StatusModel) []StatusResponse {
	var serialized []StatusResponse
	for _, status := range statuses {
		serialized = append(serialized, s.ModelToResponse(status))
	}
	return serialized
}
