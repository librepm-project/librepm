package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Transition struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	SourceStatusID uuid.UUID `json:"source_status_id"`
	TargetStatusID uuid.UUID `json:"target_status_id"`
}

type TransitionSerializer struct{}

func (s TransitionSerializer) SerializeTransition(transition domain.TransitionModel) Transition {
	return Transition{
		ID:             transition.ID,
		Name:           transition.Name,
		SourceStatusID: transition.SourceStatusID,
		TargetStatusID: transition.TargetStatusID,
	}
}

func (s TransitionSerializer) SerializeTransitions(transitions []domain.TransitionModel) []Transition {
	var serialized []Transition
	for _, transition := range transitions {
		serialized = append(serialized, s.SerializeTransition(transition))
	}
	return serialized
}
