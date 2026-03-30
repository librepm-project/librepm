package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockDashboardRepository struct {
	AllFunc        func() (*[]domain.DashboardModel, error)
	FindByIDFunc   func(id uuid.UUID) (*domain.DashboardModel, error)
	FindByNameFunc func(name string) (*domain.DashboardModel, error)
	CreateFunc     func(dashboard *domain.DashboardModel) error
	UpdateFunc     func(id uuid.UUID, dashboard *domain.DashboardModel) error
	DestroyFunc    func(id uuid.UUID) error
}

func (m *MockDashboardRepository) All() (*[]domain.DashboardModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.DashboardModel{}, nil
}

func (m *MockDashboardRepository) FindByID(id uuid.UUID) (*domain.DashboardModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.DashboardModel{}, nil
}

func (m *MockDashboardRepository) FindByName(name string) (*domain.DashboardModel, error) {
	if m.FindByNameFunc != nil {
		return m.FindByNameFunc(name)
	}
	return &domain.DashboardModel{}, nil
}

func (m *MockDashboardRepository) Create(dashboard *domain.DashboardModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(dashboard)
	}
	return nil
}

func (m *MockDashboardRepository) Update(id uuid.UUID, dashboard *domain.DashboardModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, dashboard)
	}
	return nil
}

func (m *MockDashboardRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
