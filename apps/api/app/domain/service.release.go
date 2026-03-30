package domain

import (
	"github.com/google/uuid"
)

type ReleaseServiceInterface interface {
	All() (*[]ReleaseModel, error)
	Show(id uuid.UUID) (*ReleaseModel, error)
	Create(release *ReleaseModel) error
	Update(id uuid.UUID, release *ReleaseModel) error
	Destroy(id uuid.UUID) error
}

type ReleaseService struct {
	ReleaseRepository ReleaseRepositoryInterface
}

func (s ReleaseService) All() (*[]ReleaseModel, error) {
	return s.ReleaseRepository.All()
}

func (s ReleaseService) Show(id uuid.UUID) (*ReleaseModel, error) {
	return s.ReleaseRepository.FindByID(id)
}

func (s ReleaseService) Create(release *ReleaseModel) error {
	return s.ReleaseRepository.Create(release)
}

func (s ReleaseService) Update(id uuid.UUID, release *ReleaseModel) error {
	return s.ReleaseRepository.Update(id, release)
}

func (s ReleaseService) Destroy(id uuid.UUID) error {
	return s.ReleaseRepository.Destroy(id)
}
