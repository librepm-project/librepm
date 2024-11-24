package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Dashboard struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DashboardSerializer struct{}

func (s DashboardSerializer) SerializeDashboard(dashboard domain.DashboardModel) Dashboard {
	return Dashboard{
		ID:   dashboard.ID,
		Name: dashboard.Name,
	}
}

func (s DashboardSerializer) SerializeDashboards(dashboards []domain.DashboardModel) []Dashboard {
	var serialized []Dashboard
	for _, dashboard := range dashboards {
		serialized = append(serialized, s.SerializeDashboard(dashboard))
	}
	return serialized
}
