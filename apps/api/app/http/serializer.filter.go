package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type FilterRequest struct {
	Name string `json:"name"`
}

type FilterResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type FilterSerializer struct{}

func (s FilterSerializer) RequestToModel(filter_request FilterRequest) domain.FilterModel {
	return domain.FilterModel{
		Name: filter_request.Name,
	}
}

func (s FilterSerializer) ModelToResponse(filter domain.FilterModel) FilterResponse {
	return FilterResponse{
		ID:   filter.ID,
		Name: filter.Name,
	}
}

func (s FilterSerializer) ModelToResponseMany(filters []domain.FilterModel) []FilterResponse {
	serialized := []FilterResponse{}
	for _, filter := range filters {
		serialized = append(serialized, s.ModelToResponse(filter))
	}
	return serialized
}
