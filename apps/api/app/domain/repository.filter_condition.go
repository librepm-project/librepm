package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterConditionRepositoryInterface interface {
	All() (*[]FilterConditionModel, error)
	FindByID(filter_condition_id uuid.UUID) (*FilterConditionModel, error)
	Create(filter_condition *FilterConditionModel) error
	Update(filter_condition_id uuid.UUID, filter_condition *FilterConditionModel) error
	Destroy(filter_condition_id uuid.UUID) error
}

type FilterConditionRepository struct {
	DB *gorm.DB
}

func (r FilterConditionRepository) All() (*[]FilterConditionModel, error) {
	var filter_conditions []FilterConditionModel
	var err error
	query := r.DB.Select("filter_condition.*")

	if err := query.Find(&filter_conditions).Error; err != nil {
		fmt.Println(err)
	}
	return &filter_conditions, err
}

func (r FilterConditionRepository) FindByID(filter_condition_id uuid.UUID) (*FilterConditionModel, error) {
	var filter_condition FilterConditionModel
	query := r.DB.Model(FilterConditionModel{ID: filter_condition_id}).Scan(&filter_condition)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &filter_condition, query.Error
}

func (r FilterConditionRepository) Create(filter_condition *FilterConditionModel) error {
	query := r.DB.Create(&filter_condition)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r FilterConditionRepository) Update(filter_condition_id uuid.UUID, filter_condition *FilterConditionModel) error {
	query := r.DB.Model(FilterConditionModel{}).Where("id", filter_condition_id).Updates(&filter_condition)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r FilterConditionRepository) Destroy(filter_condition_id uuid.UUID) error {
	query := r.DB.Model(FilterConditionModel{}).Delete(FilterConditionModel{}, filter_condition_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
