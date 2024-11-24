package domain

import (
	"github.com/google/uuid"
)

type FilterServiceInterface interface {
	All() *[]FilterModel
	Show(filter_id uuid.UUID) *FilterModel
	Create(filter *FilterModel)
	Update(filter_id uuid.UUID, filter *FilterModel)
	Destroy(filter_id uuid.UUID)
}

type FilterService struct {
	FilterRepository FilterRepositoryInterface
}

func (s FilterService) All() *[]FilterModel {
	filters := s.FilterRepository.All()
	return filters
}

func (s FilterService) Show(filter_id uuid.UUID) *FilterModel {
	return s.FilterRepository.FindByID(filter_id)
}

func (s FilterService) Create(filter *FilterModel) {
	s.FilterRepository.Create(filter)
}

func (s FilterService) Update(filter_id uuid.UUID, filter *FilterModel) {
	s.FilterRepository.Update(filter_id, filter)
}

func (s FilterService) Destroy(filter_id uuid.UUID) {
	s.FilterRepository.Destroy(filter_id)
}
