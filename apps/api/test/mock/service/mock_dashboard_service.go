package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockDashboardService struct {
	AllFunc     func() (*[]domain.DashboardModel, error)
	ShowFunc    func(dashboardID uuid.UUID) (*domain.DashboardModel, error)
	CreateFunc  func(dashboard *domain.DashboardModel) error
	UpdateFunc  func(dashboardID uuid.UUID, dashboard *domain.DashboardModel) error
	DestroyFunc func(dashboardID uuid.UUID) error
}

func (m *MockDashboardService) All() (*[]domain.DashboardModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.DashboardModel{}, nil
}

func (m *MockDashboardService) Show(dashboardID uuid.UUID) (*domain.DashboardModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(dashboardID)
	}
	return &domain.DashboardModel{}, nil
}

func (m *MockDashboardService) Create(dashboard *domain.DashboardModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(dashboard)
	}
	return nil
}

func (m *MockDashboardService) Update(dashboardID uuid.UUID, dashboard *domain.DashboardModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(dashboardID, dashboard)
	}
	return nil
}

func (m *MockDashboardService) Destroy(dashboardID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(dashboardID)
	}
	return nil
}
