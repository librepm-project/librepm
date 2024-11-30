package domain

import (
	"github.com/google/uuid"
)

type TransitionServiceInterface interface {
	All() *[]TransitionModel
	Show(transition_id uuid.UUID) *TransitionModel
	Create(transition *TransitionModel)
	Update(transition_id uuid.UUID, transition *TransitionModel)
	Destroy(transition_id uuid.UUID)
}

type TransitionService struct {
	TransitionRepository TransitionRepositoryInterface
}

func (s TransitionService) All() *[]TransitionModel {
	transitions := s.TransitionRepository.All()
	return transitions
}

func (s TransitionService) Show(transition_id uuid.UUID) *TransitionModel {
	return s.TransitionRepository.FindByID(transition_id)
}

func (s TransitionService) Create(transition *TransitionModel) {
	s.TransitionRepository.Create(transition)
}

func (s TransitionService) Update(transition_id uuid.UUID, transition *TransitionModel) {
	s.TransitionRepository.Update(transition_id, transition)
}

func (s TransitionService) Destroy(transition_id uuid.UUID) {
	s.TransitionRepository.Destroy(transition_id)
}
