package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockDashboardWidgetService struct {
	AllByDashboardFunc    func(dashboardID uuid.UUID) (*[]domain.DashboardWidgetModel, error)
	CreateFunc            func(widget *domain.DashboardWidgetModel) error
	UpdateFunc            func(widgetID uuid.UUID, fields map[string]interface{}) error
	BatchUpdateWeightsFunc func(updates []domain.WeightUpdate) error
	DestroyFunc           func(widgetID uuid.UUID) error
}

func (m *MockDashboardWidgetService) AllByDashboard(dashboardID uuid.UUID) (*[]domain.DashboardWidgetModel, error) {
	if m.AllByDashboardFunc != nil {
		return m.AllByDashboardFunc(dashboardID)
	}
	return &[]domain.DashboardWidgetModel{}, nil
}

func (m *MockDashboardWidgetService) Create(widget *domain.DashboardWidgetModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(widget)
	}
	return nil
}

func (m *MockDashboardWidgetService) Update(widgetID uuid.UUID, fields map[string]interface{}) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(widgetID, fields)
	}
	return nil
}

func (m *MockDashboardWidgetService) BatchUpdateWeights(updates []domain.WeightUpdate) error {
	if m.BatchUpdateWeightsFunc != nil {
		return m.BatchUpdateWeightsFunc(updates)
	}
	return nil
}

func (m *MockDashboardWidgetService) Destroy(widgetID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(widgetID)
	}
	return nil
}
