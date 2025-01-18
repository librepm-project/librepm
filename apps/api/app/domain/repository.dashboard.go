package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardRepositoryInterface interface {
	All() (*[]DashboardModel, error)
	FindByID(dashboard_id uuid.UUID) (*DashboardModel, error)
	Create(dashboard *DashboardModel) error
	Update(dashboard_id uuid.UUID, dashboard *DashboardModel) error
	Destroy(dashboard_id uuid.UUID) error
}

type DashboardRepository struct {
	DB *gorm.DB
}

func (r DashboardRepository) All() (*[]DashboardModel, error) {
	var dashboards []DashboardModel
	var err error
	query := r.DB.Select("dashboard.*")

	if err := query.Find(&dashboards).Error; err != nil {
		fmt.Println(err)
	}
	return &dashboards, err
}

func (r DashboardRepository) FindByID(dashboard_id uuid.UUID) (*DashboardModel, error) {
	var dashboard DashboardModel
	query := r.DB.Model(DashboardModel{ID: dashboard_id}).Scan(&dashboard)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &dashboard, query.Error
}

func (r DashboardRepository) Create(dashboard *DashboardModel) error {
	query := r.DB.Create(&dashboard)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r DashboardRepository) Update(dashboard_id uuid.UUID, dashboard *DashboardModel) error {
	query := r.DB.Model(DashboardModel{}).Where("id", dashboard_id).Updates(&dashboard)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r DashboardRepository) Destroy(dashboard_id uuid.UUID) error {
	query := r.DB.Model(DashboardModel{}).Delete(DashboardModel{}, dashboard_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
