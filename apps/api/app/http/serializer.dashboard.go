package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type DashboardRequest struct {
	Name string `json:"name"`
}

type DashboardResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DashboardSerializer struct{}

func (s DashboardSerializer) RequestToModel(dashboard_request DashboardRequest) domain.DashboardModel {
	return domain.DashboardModel{
		Name: dashboard_request.Name,
	}
}

func (s DashboardSerializer) ModelToResponse(dashboard domain.DashboardModel) DashboardResponse {
	return DashboardResponse{
		ID:   dashboard.ID,
		Name: dashboard.Name,
	}
}

func (s DashboardSerializer) ModelToResponseMany(dashboards []domain.DashboardModel) []DashboardResponse {
	var serialized []DashboardResponse
	for _, dashboard := range dashboards {
		serialized = append(serialized, s.ModelToResponse(dashboard))
	}
	return serialized
}
