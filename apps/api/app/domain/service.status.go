package domain

import (
	"github.com/google/uuid"
)

type StatusServiceInterface interface {
	All() *[]StatusModel
	Show(status_id uuid.UUID) *StatusModel
	Create(status *StatusModel)
	Update(status_id uuid.UUID, status *StatusModel)
	Destroy(status_id uuid.UUID)
}

type StatusService struct {
	StatusRepository StatusRepositoryInterface
}

func (s StatusService) All() *[]StatusModel {
	statuses := s.StatusRepository.All()
	return statuses
}

func (s StatusService) Show(status_id uuid.UUID) *StatusModel {
	return s.StatusRepository.FindByID(status_id)
}

func (s StatusService) Create(status *StatusModel) {
	s.StatusRepository.Create(status)
}

func (s StatusService) Update(status_id uuid.UUID, status *StatusModel) {
	s.StatusRepository.Update(status_id, status)
}

func (s StatusService) Destroy(status_id uuid.UUID) {
	s.StatusRepository.Destroy(status_id)
}
