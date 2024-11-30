package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterConditionRepositoryInterface interface {
	All() *[]FilterConditionModel
	FindByID(filter_condition_id uuid.UUID) *FilterConditionModel
	Create(filter_condition *FilterConditionModel)
	Update(filter_condition_id uuid.UUID, filter_condition *FilterConditionModel)
	Destroy(filter_condition_id uuid.UUID)
}

type FilterConditionRepository struct {
	DB *gorm.DB
}

func (r FilterConditionRepository) All() *[]FilterConditionModel {
	var filter_conditions []FilterConditionModel
	query := r.DB.Select("filter_condition.*")

	if err := query.Find(&filter_conditions).Error; err != nil {
		fmt.Println(err)
	}
	return &filter_conditions
}

func (r FilterConditionRepository) FindByID(filter_condition_id uuid.UUID) *FilterConditionModel {
	var filter_condition FilterConditionModel
	fmt.Println(filter_condition_id)
	err := r.DB.Model(FilterConditionModel{ID: filter_condition_id}).Scan(&filter_condition)

	if err != nil {
		fmt.Println(err)
	}
	return &filter_condition
}

func (r FilterConditionRepository) Create(filter_condition *FilterConditionModel) {
	r.DB.Create(&filter_condition)
}

func (r FilterConditionRepository) Update(filter_condition_id uuid.UUID, filter_condition *FilterConditionModel) {
	r.DB.Model(FilterConditionModel{}).Where("id", filter_condition_id).Updates(&filter_condition)
}

func (r FilterConditionRepository) Destroy(filter_condition_id uuid.UUID) {
	r.DB.Model(FilterConditionModel{}).Delete(FilterConditionModel{}, filter_condition_id)
}
