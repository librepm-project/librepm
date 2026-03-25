package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterRepositoryInterface interface {
	All() (*[]FilterModel, error)
	FindByID(filter_id uuid.UUID) (*FilterModel, error)
	Create(filter *FilterModel) error
	Update(filter_id uuid.UUID, filter *FilterModel) error
	Destroy(filter_id uuid.UUID) error
}

type FilterRepository struct {
	DB *gorm.DB
}

func (r FilterRepository) All() (*[]FilterModel, error) {
	var filters []FilterModel
	if err := r.DB.Preload("FilterConditions").Find(&filters).Error; err != nil {
		slog.Error("failed to fetch filters", "error", err)
		return nil, err
	}
	return &filters, nil
}

func (r FilterRepository) FindByID(filter_id uuid.UUID) (*FilterModel, error) {
	var filter FilterModel
	query := r.DB.Preload("FilterConditions").First(&filter, filter_id)
	if query.Error != nil {
		slog.Error("failed to find filter by id", "error", query.Error)
	}
	return &filter, query.Error
}

func (r FilterRepository) Create(filter *FilterModel) error {
	query := r.DB.Create(&filter)
	if query.Error != nil {
		slog.Error("failed to create filter", "error", query.Error)
	}
	return query.Error
}

func (r FilterRepository) Update(filter_id uuid.UUID, filter *FilterModel) error {
	if err := r.DB.Where("filter_id = ?", filter_id).Delete(&FilterConditionModel{}).Error; err != nil {
		return err
	}
	if err := r.DB.Model(FilterModel{}).Where("id = ?", filter_id).Updates(map[string]interface{}{
		"name":        filter.Name,
		"description": filter.Description,
	}).Error; err != nil {
		return err
	}
	if len(filter.FilterConditions) > 0 {
		for i := range filter.FilterConditions {
			filter.FilterConditions[i].FilterID = filter_id
		}
		if err := r.DB.Create(&filter.FilterConditions).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r FilterRepository) Destroy(filter_id uuid.UUID) error {
	r.DB.Where("filter_id = ?", filter_id).Delete(&FilterConditionModel{})
	query := r.DB.Delete(FilterModel{}, filter_id)
	if query.Error != nil {
		slog.Error("failed to destroy filter", "error", query.Error)
	}
	return query.Error
}
