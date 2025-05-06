package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type StatusRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type StatusResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Color string    `json:"color"`
}

type StatusSerializer struct{}

func (s StatusSerializer) RequestToModel(status_request StatusRequest) domain.StatusModel {
	return domain.StatusModel{
		Name:  status_request.Name,
		Color: status_request.Color,
	}
}

func (s StatusSerializer) ModelToResponse(status domain.StatusModel) StatusResponse {
	return StatusResponse{
		ID:    status.ID,
		Name:  status.Name,
		Color: status.Color,
	}
}

func (s StatusSerializer) ModelToResponseMany(statuses []domain.StatusModel) []StatusResponse {
	serialized := []StatusResponse{}
	for _, status := range statuses {
		serialized = append(serialized, s.ModelToResponse(status))
	}
	return serialized
}
