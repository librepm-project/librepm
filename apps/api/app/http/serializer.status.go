package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Status struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type StatusSerializer struct{}

func (s StatusSerializer) SerializeStatus(status domain.StatusModel) Status {
	return Status{
		ID:   status.ID,
		Name: status.Name,
	}
}

func (s StatusSerializer) SerializeStatuss(statuses []domain.StatusModel) []Status {
	var serialized []Status
	for _, status := range statuses {
		serialized = append(serialized, s.SerializeStatus(status))
	}
	return serialized
}
