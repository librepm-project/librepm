package domain

import (
	"github.com/google/uuid"
)

type DashboardWidgetServiceInterface interface {
	AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error)
	Create(widget *DashboardWidgetModel) error
	Destroy(widget_id uuid.UUID) error
}

type DashboardWidgetService struct {
	DashboardWidgetRepository DashboardWidgetRepositoryInterface
}

func (s DashboardWidgetService) AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error) {
	return s.DashboardWidgetRepository.AllByDashboard(dashboard_id)
}

func (s DashboardWidgetService) Create(widget *DashboardWidgetModel) error {
	return s.DashboardWidgetRepository.Create(widget)
}

func (s DashboardWidgetService) Destroy(widget_id uuid.UUID) error {
	return s.DashboardWidgetRepository.Destroy(widget_id)
}
