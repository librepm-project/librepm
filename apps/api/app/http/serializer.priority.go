package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type PriorityRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type PriorityResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Color string    `json:"color"`
}

type PrioritySerializer struct{}

func (s PrioritySerializer) RequestToModel(r PriorityRequest) domain.PriorityModel {
	return domain.PriorityModel{
		Name:  r.Name,
		Color: r.Color,
	}
}

func (s PrioritySerializer) ModelToResponse(priority domain.PriorityModel) PriorityResponse {
	return PriorityResponse{
		ID:    priority.ID,
		Name:  priority.Name,
		Color: priority.Color,
	}
}

func (s PrioritySerializer) ModelToResponseMany(priorities []domain.PriorityModel) []PriorityResponse {
	serialized := []PriorityResponse{}
	for _, p := range priorities {
		serialized = append(serialized, s.ModelToResponse(p))
	}
	return serialized
}
