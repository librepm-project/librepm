package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardWidgetRepositoryInterface interface {
	AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error)
	Create(widget *DashboardWidgetModel) error
	Destroy(widget_id uuid.UUID) error
}

type DashboardWidgetRepository struct {
	DB *gorm.DB
}

func (r DashboardWidgetRepository) AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error) {
	var widgets []DashboardWidgetModel
	if err := r.DB.
		Preload("Filter.FilterConditions").
		Where("dashboard_id = ?", dashboard_id).
		Find(&widgets).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &widgets, nil
}

func (r DashboardWidgetRepository) Create(widget *DashboardWidgetModel) error {
	query := r.DB.Create(&widget)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query.Error
}

func (r DashboardWidgetRepository) Destroy(widget_id uuid.UUID) error {
	query := r.DB.Delete(&DashboardWidgetModel{}, widget_id)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query.Error
}
