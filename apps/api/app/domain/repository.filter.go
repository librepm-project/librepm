package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterRepositoryInterface interface {
	All() *[]FilterModel
	FindByID(filter_id uuid.UUID) *FilterModel
	Create(filter *FilterModel)
	Update(filter_id uuid.UUID, filter *FilterModel)
	Destroy(filter_id uuid.UUID)
}

type FilterRepository struct {
	DB *gorm.DB
}

func (r FilterRepository) All() *[]FilterModel {
	var filters []FilterModel
	query := r.DB.Select("filter.*")

	if err := query.Find(&filters).Error; err != nil {
		fmt.Println(err)
	}
	return &filters
}

func (r FilterRepository) FindByID(filter_id uuid.UUID) *FilterModel {
	var filter FilterModel
	fmt.Println(filter_id)
	err := r.DB.Model(FilterModel{ID: filter_id}).Scan(&filter)

	if err != nil {
		fmt.Println(err)
	}
	return &filter
}

func (r FilterRepository) Create(filter *FilterModel) {
	r.DB.Create(&filter)
}

func (r FilterRepository) Update(filter_id uuid.UUID, filter *FilterModel) {
	r.DB.Model(FilterModel{}).Where("id", filter_id).Updates(&filter)
}

func (r FilterRepository) Destroy(filter_id uuid.UUID) {
	r.DB.Model(FilterModel{}).Delete(FilterModel{}, filter_id)
}
