package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardWidgetRepositoryInterface interface {
	AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error)
	Create(widget *DashboardWidgetModel) error
	Update(widget_id uuid.UUID, fields map[string]interface{}) error
	BatchUpdateWeights(updates []WeightUpdate) error
	Destroy(widget_id uuid.UUID) error
}

type WeightUpdate struct {
	ID     uuid.UUID
	Weight int
}

type DashboardWidgetRepository struct {
	DB *gorm.DB
}

func (r DashboardWidgetRepository) AllByDashboard(dashboard_id uuid.UUID) (*[]DashboardWidgetModel, error) {
	var widgets []DashboardWidgetModel
	if err := r.DB.
		Preload("Filter.FilterConditions").
		Where("dashboard_id = ?", dashboard_id).
		Order("weight ASC").
		Find(&widgets).Error; err != nil {
		slog.Error("failed to fetch dashboard widgets", "error", err)
		return nil, err
	}
	return &widgets, nil
}

func (r DashboardWidgetRepository) Create(widget *DashboardWidgetModel) error {
	query := r.DB.Create(&widget)
	if query.Error != nil {
		slog.Error("failed to create dashboard widget", "error", query.Error)
	}
	return query.Error
}

func (r DashboardWidgetRepository) Update(widget_id uuid.UUID, fields map[string]interface{}) error {
	query := r.DB.Model(&DashboardWidgetModel{}).Where("id = ?", widget_id).Updates(fields)
	if query.Error != nil {
		slog.Error("failed to update dashboard widget", "error", query.Error)
	}
	return query.Error
}

func (r DashboardWidgetRepository) BatchUpdateWeights(updates []WeightUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for _, u := range updates {
			if err := tx.Model(&DashboardWidgetModel{}).Where("id = ?", u.ID).Update("weight", u.Weight).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r DashboardWidgetRepository) Destroy(widget_id uuid.UUID) error {
	query := r.DB.Delete(&DashboardWidgetModel{}, widget_id)
	if query.Error != nil {
		slog.Error("failed to destroy dashboard widget", "error", query.Error)
	}
	return query.Error
}
