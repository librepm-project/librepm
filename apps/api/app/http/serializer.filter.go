package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type FilterConditionRequest struct {
	Field string `json:"field"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

type FilterConditionResponse struct {
	ID    uuid.UUID `json:"id"`
	Field string    `json:"field"`
	Op    string    `json:"op"`
	Value string    `json:"value"`
}

type FilterRequest struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Public      bool                     `json:"public"`
	ColumnList  string                   `json:"columnList"`
	GroupBy     string                   `json:"groupBy"`
	Conditions  []FilterConditionRequest `json:"conditions"`
}

type FilterResponse struct {
	ID          uuid.UUID                 `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Public      bool                      `json:"public"`
	ColumnList  string                    `json:"columnList"`
	GroupBy     string                    `json:"groupBy"`
	Conditions  []FilterConditionResponse `json:"conditions"`
}

type FilterSerializer struct{}

func (s FilterSerializer) RequestToModel(filter_request FilterRequest) domain.FilterModel {
	conditions := []domain.FilterConditionModel{}
	for _, c := range filter_request.Conditions {
		conditions = append(conditions, domain.FilterConditionModel{
			Field: c.Field,
			Op:    c.Op,
			Value: c.Value,
		})
	}
	return domain.FilterModel{
		Name:             filter_request.Name,
		Description:      filter_request.Description,
		Public:           filter_request.Public,
		ColumnList:       filter_request.ColumnList,
		GroupBy:          filter_request.GroupBy,
		FilterConditions: conditions,
	}
}

func (s FilterSerializer) ModelToResponse(filter domain.FilterModel) FilterResponse {
	conditions := []FilterConditionResponse{}
	for _, c := range filter.FilterConditions {
		conditions = append(conditions, FilterConditionResponse{
			ID:    c.ID,
			Field: c.Field,
			Op:    c.Op,
			Value: c.Value,
		})
	}
	return FilterResponse{
		ID:          filter.ID,
		Name:        filter.Name,
		Description: filter.Description,
		Public:      filter.Public,
		ColumnList:  filter.ColumnList,
		GroupBy:     filter.GroupBy,
		Conditions:  conditions,
	}
}

func (s FilterSerializer) ModelToResponseMany(filters []domain.FilterModel) []FilterResponse {
	serialized := []FilterResponse{}
	for _, filter := range filters {
		serialized = append(serialized, s.ModelToResponse(filter))
	}
	return serialized
}
