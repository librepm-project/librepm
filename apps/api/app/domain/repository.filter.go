package domain

import (
	"fmt"

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
	var err error
	query := r.DB.Select("filter.*")

	if err := query.Find(&filters).Error; err != nil {
		fmt.Println(err)
	}
	return &filters, err
}

func (r FilterRepository) FindByID(filter_id uuid.UUID) (*FilterModel, error) {
	var filter FilterModel
	query := r.DB.Model(FilterModel{ID: filter_id}).Scan(&filter)

	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return &filter, query.Error
}

func (r FilterRepository) Create(filter *FilterModel) error {
	query := r.DB.Create(&filter)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r FilterRepository) Update(filter_id uuid.UUID, filter *FilterModel) error {
	query := r.DB.Model(FilterModel{}).Where("id", filter_id).Updates(&filter)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r FilterRepository) Destroy(filter_id uuid.UUID) error {
	query := r.DB.Model(FilterModel{}).Delete(FilterModel{}, filter_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
