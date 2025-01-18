package domain

import (
	"github.com/google/uuid"
)

type StatusServiceInterface interface {
	All() (*[]StatusModel, error)
	Show(status_id uuid.UUID) (*StatusModel, error)
	Create(status *StatusModel) error
	Update(status_id uuid.UUID, status *StatusModel) error
	Destroy(status_id uuid.UUID) error
}

type StatusService struct {
	StatusRepository StatusRepositoryInterface
}

func (s StatusService) All() (*[]StatusModel, error) {
	statuses, err := s.StatusRepository.All()
	return statuses, err
}

func (s StatusService) Show(status_id uuid.UUID) (*StatusModel, error) {
	return s.StatusRepository.FindByID(status_id)
}

func (s StatusService) Create(status *StatusModel) error {
	return s.StatusRepository.Create(status)
}

func (s StatusService) Update(status_id uuid.UUID, status *StatusModel) error {
	return s.StatusRepository.Update(status_id, status)
}

func (s StatusService) Destroy(status_id uuid.UUID) error {
	return s.StatusRepository.Destroy(status_id)
}
