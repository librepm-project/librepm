package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type TransitionRequest struct {
	Name           string    `json:"name"`
	SourceStatusID uuid.UUID `json:"sourceStatusId"`
	TargetStatusID uuid.UUID `json:"targetStatusId"`
}

type TransitionResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SourceStatusID uuid.UUID `json:"sourceStatusId"`
	TargetStatusID uuid.UUID `json:"targetStatusId"`
}

type TransitionSerializer struct{}

func (s TransitionSerializer) RequestToModel(transition_request TransitionRequest) domain.TransitionModel {
	return domain.TransitionModel{
		Name:           transition_request.Name,
		SourceStatusID: transition_request.SourceStatusID,
		TargetStatusID: transition_request.TargetStatusID,
	}
}

func (s TransitionSerializer) ModelToResponse(transition domain.TransitionModel) TransitionResponse {
	return TransitionResponse{
		ID:             transition.ID,
		Name:           transition.Name,
		SourceStatusID: transition.SourceStatusID,
		TargetStatusID: transition.TargetStatusID,
	}
}

func (s TransitionSerializer) ModelToResponseMany(transitions []domain.TransitionModel) []TransitionResponse {
	serialized := []TransitionResponse{}
	for _, transition := range transitions {
		serialized = append(serialized, s.ModelToResponse(transition))
	}
	return serialized
}
