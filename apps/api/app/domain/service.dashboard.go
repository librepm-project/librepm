package domain

import (
	"github.com/google/uuid"
)

type DashboardServiceInterface interface {
	All() (*[]DashboardModel, error)
	Show(dashboard_id uuid.UUID) (*DashboardModel, error)
	Create(dashboard *DashboardModel) error
	Update(dashboard_id uuid.UUID, dashboard *DashboardModel) error
	Destroy(dashboard_id uuid.UUID) error
}

type DashboardService struct {
	DashboardRepository DashboardRepositoryInterface
}

func (s DashboardService) All() (*[]DashboardModel, error) {
	dashboards, err := s.DashboardRepository.All()
	return dashboards, err
}

func (s DashboardService) Show(dashboard_id uuid.UUID) (*DashboardModel, error) {
	return s.DashboardRepository.FindByID(dashboard_id)
}

func (s DashboardService) Create(dashboard *DashboardModel) error {
	return s.DashboardRepository.Create(dashboard)
}

func (s DashboardService) Update(dashboard_id uuid.UUID, dashboard *DashboardModel) error {
	return s.DashboardRepository.Update(dashboard_id, dashboard)
}

func (s DashboardService) Destroy(dashboard_id uuid.UUID) error {
	return s.DashboardRepository.Destroy(dashboard_id)
}
