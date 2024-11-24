package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardRepositoryInterface interface {
	All() *[]DashboardModel
	FindByID(dashboard_id uuid.UUID) *DashboardModel
	Create(dashboard *DashboardModel)
	Update(dashboard_id uuid.UUID, dashboard *DashboardModel)
	Destroy(dashboard_id uuid.UUID)
}

type DashboardRepository struct {
	DB *gorm.DB
}

func (r DashboardRepository) All() *[]DashboardModel {
	var dashboards []DashboardModel
	query := r.DB.Select("dashboard.*")

	if err := query.Find(&dashboards).Error; err != nil {
		fmt.Println(err)
	}
	return &dashboards
}

func (r DashboardRepository) FindByID(dashboard_id uuid.UUID) *DashboardModel {
	var dashboard DashboardModel
	fmt.Println(dashboard_id)
	err := r.DB.Model(DashboardModel{ID: dashboard_id}).Scan(&dashboard)

	if err != nil {
		fmt.Println(err)
	}
	return &dashboard
}

func (r DashboardRepository) Create(dashboard *DashboardModel) {
	r.DB.Create(&dashboard)
}

func (r DashboardRepository) Update(dashboard_id uuid.UUID, dashboard *DashboardModel) {
	r.DB.Model(DashboardModel{}).Where("id", dashboard_id).Updates(&dashboard)
}

func (r DashboardRepository) Destroy(dashboard_id uuid.UUID) {
	r.DB.Model(DashboardModel{}).Delete(DashboardModel{}, dashboard_id)
}
