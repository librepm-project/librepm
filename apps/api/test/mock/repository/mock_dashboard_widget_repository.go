package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockDashboardWidgetRepository struct {
	AllByDashboardFunc    func(dashboardID uuid.UUID) (*[]domain.DashboardWidgetModel, error)
	CreateFunc            func(widget *domain.DashboardWidgetModel) error
	UpdateFunc            func(id uuid.UUID, fields map[string]interface{}) error
	BatchUpdateWeightsFunc func(updates []domain.WeightUpdate) error
	DestroyFunc           func(id uuid.UUID) error
}

func (m *MockDashboardWidgetRepository) AllByDashboard(dashboardID uuid.UUID) (*[]domain.DashboardWidgetModel, error) {
	if m.AllByDashboardFunc != nil {
		return m.AllByDashboardFunc(dashboardID)
	}
	return &[]domain.DashboardWidgetModel{}, nil
}

func (m *MockDashboardWidgetRepository) Create(widget *domain.DashboardWidgetModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(widget)
	}
	return nil
}

func (m *MockDashboardWidgetRepository) Update(id uuid.UUID, fields map[string]interface{}) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, fields)
	}
	return nil
}

func (m *MockDashboardWidgetRepository) BatchUpdateWeights(updates []domain.WeightUpdate) error {
	if m.BatchUpdateWeightsFunc != nil {
		return m.BatchUpdateWeightsFunc(updates)
	}
	return nil
}

func (m *MockDashboardWidgetRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
