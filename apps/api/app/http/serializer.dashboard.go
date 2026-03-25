package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type DashboardRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

type DashboardResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Public      bool      `json:"public"`
}

type DashboardSerializer struct{}

func (s DashboardSerializer) RequestToModel(dashboard_request DashboardRequest) domain.DashboardModel {
	return domain.DashboardModel{
		Name:        dashboard_request.Name,
		Description: dashboard_request.Description,
		Public:      dashboard_request.Public,
	}
}

func (s DashboardSerializer) ModelToResponse(dashboard domain.DashboardModel) DashboardResponse {
	return DashboardResponse{
		ID:          dashboard.ID,
		Name:        dashboard.Name,
		Description: dashboard.Description,
		Public:      dashboard.Public,
	}
}

func (s DashboardSerializer) ModelToResponseMany(dashboards []domain.DashboardModel) []DashboardResponse {
	serialized := []DashboardResponse{}
	for _, dashboard := range dashboards {
		serialized = append(serialized, s.ModelToResponse(dashboard))
	}
	return serialized
}
