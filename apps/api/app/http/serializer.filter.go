package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Filter struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type FilterSerializer struct{}

func (s FilterSerializer) SerializeFilter(filter domain.FilterModel) Filter {
	return Filter{
		ID:   filter.ID,
		Name: filter.Name,
	}
}

func (s FilterSerializer) SerializeFilters(filters []domain.FilterModel) []Filter {
	var serialized []Filter
	for _, filter := range filters {
		serialized = append(serialized, s.SerializeFilter(filter))
	}
	return serialized
}
