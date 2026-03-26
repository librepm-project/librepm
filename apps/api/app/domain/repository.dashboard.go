package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardRepositoryInterface interface {
	All() (*[]DashboardModel, error)
	FindByID(dashboard_id uuid.UUID) (*DashboardModel, error)
	FindByName(name string) (*DashboardModel, error)
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
	query := r.DB.Preload("Widgets.Filter.FilterConditions").Select("dashboard.*")

	if err := query.Find(&dashboards).Error; err != nil {
		slog.Error("failed to fetch dashboards", "error", err)
	}
	return &dashboards, err
}

func (r DashboardRepository) FindByID(dashboard_id uuid.UUID) (*DashboardModel, error) {
	var dashboard DashboardModel
	query := r.DB.Preload("Widgets.Filter.FilterConditions").First(&dashboard, dashboard_id)

	if query.Error != nil {
		slog.Error("failed to find dashboard by id", "error", query.Error)
	}
	return &dashboard, query.Error
}

func (r DashboardRepository) FindByName(name string) (*DashboardModel, error) {
	var dashboard DashboardModel
	query := r.DB.Where("name = ?", name).First(&dashboard)

	if query.Error != nil {
		slog.Error("failed to find dashboard by name", "name", name, "error", query.Error)
	}
	return &dashboard, query.Error
}

func (r DashboardRepository) Create(dashboard *DashboardModel) error {
	query := r.DB.Create(&dashboard)
	if query.Error != nil {
		slog.Error("failed to create dashboard", "error", query.Error)
	}
	return query.Error
}

func (r DashboardRepository) Update(dashboard_id uuid.UUID, dashboard *DashboardModel) error {
	query := r.DB.Model(DashboardModel{}).Where("id", dashboard_id).Updates(&dashboard)
	if query.Error != nil {
		slog.Error("failed to update dashboard", "error", query.Error)
	}
	return query.Error
}

func (r DashboardRepository) Destroy(dashboard_id uuid.UUID) error {
	query := r.DB.Model(DashboardModel{}).Delete(DashboardModel{}, dashboard_id)
	if query.Error != nil {
		slog.Error("failed to destroy dashboard", "error", query.Error)
	}
	return query.Error
}
