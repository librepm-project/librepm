package domain

import (
	"github.com/google/uuid"
)

type DashboardServiceInterface interface {
	All() *[]DashboardModel
	Show(dashboard_id uuid.UUID) *DashboardModel
	Create(dashboard *DashboardModel)
	Update(dashboard_id uuid.UUID, dashboard *DashboardModel)
	Destroy(dashboard_id uuid.UUID)
}

type DashboardService struct {
	DashboardRepository DashboardRepositoryInterface
}

func (s DashboardService) All() *[]DashboardModel {
	dashboards := s.DashboardRepository.All()
	return dashboards
}

func (s DashboardService) Show(dashboard_id uuid.UUID) *DashboardModel {
	return s.DashboardRepository.FindByID(dashboard_id)
}

func (s DashboardService) Create(dashboard *DashboardModel) {
	s.DashboardRepository.Create(dashboard)
}

func (s DashboardService) Update(dashboard_id uuid.UUID, dashboard *DashboardModel) {
	s.DashboardRepository.Update(dashboard_id, dashboard)
}

func (s DashboardService) Destroy(dashboard_id uuid.UUID) {
	s.DashboardRepository.Destroy(dashboard_id)
}
