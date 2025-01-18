package domain

import (
	"github.com/google/uuid"
)

type TransitionServiceInterface interface {
	All() (*[]TransitionModel, error)
	Show(transition_id uuid.UUID) (*TransitionModel, error)
	Create(transition *TransitionModel) error
	Update(transition_id uuid.UUID, transition *TransitionModel) error
	Destroy(transition_id uuid.UUID) error
}

type TransitionService struct {
	TransitionRepository TransitionRepositoryInterface
}

func (s TransitionService) All() (*[]TransitionModel, error) {
	transitions, err := s.TransitionRepository.All()
	return transitions, err
}

func (s TransitionService) Show(transition_id uuid.UUID) (*TransitionModel, error) {
	return s.TransitionRepository.FindByID(transition_id)
}

func (s TransitionService) Create(transition *TransitionModel) error {
	return s.TransitionRepository.Create(transition)
}

func (s TransitionService) Update(transition_id uuid.UUID, transition *TransitionModel) error {
	return s.TransitionRepository.Update(transition_id, transition)
}

func (s TransitionService) Destroy(transition_id uuid.UUID) error {
	return s.TransitionRepository.Destroy(transition_id)
}
