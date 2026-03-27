package domain

import (
	"github.com/google/uuid"
)

type PriorityServiceInterface interface {
	All() (*[]PriorityModel, error)
	Show(priority_id uuid.UUID) (*PriorityModel, error)
	Create(priority *PriorityModel) error
	Update(priority_id uuid.UUID, priority *PriorityModel) error
	Destroy(priority_id uuid.UUID) error
}

type PriorityService struct {
	PriorityRepository PriorityRepositoryInterface
}

func (s PriorityService) All() (*[]PriorityModel, error) {
	return s.PriorityRepository.All()
}

func (s PriorityService) Show(priority_id uuid.UUID) (*PriorityModel, error) {
	return s.PriorityRepository.FindByID(priority_id)
}

func (s PriorityService) Create(priority *PriorityModel) error {
	return s.PriorityRepository.Create(priority)
}

func (s PriorityService) Update(priority_id uuid.UUID, priority *PriorityModel) error {
	return s.PriorityRepository.Update(priority_id, priority)
}

func (s PriorityService) Destroy(priority_id uuid.UUID) error {
	return s.PriorityRepository.Destroy(priority_id)
}
