package domain

import (
	"github.com/google/uuid"
)

type DashboardWidgetServiceInterface interface {
	AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error)
	Create(widget *DashboardWidgetModel) error
	Update(widget_id uuid.UUID, fields map[string]interface{}) error
	BatchUpdateWeights(updates []WeightUpdate) error
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

func (s DashboardWidgetService) Update(widget_id uuid.UUID, fields map[string]interface{}) error {
	return s.DashboardWidgetRepository.Update(widget_id, fields)
}

func (s DashboardWidgetService) BatchUpdateWeights(updates []WeightUpdate) error {
	return s.DashboardWidgetRepository.BatchUpdateWeights(updates)
}

func (s DashboardWidgetService) Destroy(widget_id uuid.UUID) error {
	return s.DashboardWidgetRepository.Destroy(widget_id)
}
