package domain

import (
	"github.com/google/uuid"
)

type FilterServiceInterface interface {
	All() (*[]FilterModel, error)
	Show(filter_id uuid.UUID) (*FilterModel, error)
	Create(filter *FilterModel) error
	Update(filter_id uuid.UUID, filter *FilterModel) error
	Destroy(filter_id uuid.UUID) error
}

type FilterService struct {
	FilterRepository FilterRepositoryInterface
}

func (s FilterService) All() (*[]FilterModel, error) {
	filters, err := s.FilterRepository.All()
	return filters, err
}

func (s FilterService) Show(filter_id uuid.UUID) (*FilterModel, error) {
	return s.FilterRepository.FindByID(filter_id)
}

func (s FilterService) Create(filter *FilterModel) error {
	return s.FilterRepository.Create(filter)
}

func (s FilterService) Update(filter_id uuid.UUID, filter *FilterModel) error {
	return s.FilterRepository.Update(filter_id, filter)
}

func (s FilterService) Destroy(filter_id uuid.UUID) error {
	return s.FilterRepository.Destroy(filter_id)
}
