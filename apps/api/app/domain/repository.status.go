package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusRepositoryInterface interface {
	All() (*[]StatusModel, error)
	FindByID(status_id uuid.UUID) (*StatusModel, error)
	FindByName(name string) (*StatusModel, error)
	Create(status *StatusModel) error
	Update(status_id uuid.UUID, status *StatusModel) error
	Destroy(status_id uuid.UUID) error
}

type StatusRepository struct {
	DB *gorm.DB
}

func (r StatusRepository) All() (*[]StatusModel, error) {
	var statuses []StatusModel
	var err error
	query := r.DB.Select("status.*")

	if err := query.Find(&statuses).Error; err != nil {
		fmt.Println(err)
	}
	return &statuses, err
}

func (r StatusRepository) FindByID(status_id uuid.UUID) (*StatusModel, error) {
	var status StatusModel
	query := r.DB.Model(StatusModel{ID: status_id}).Scan(&status)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &status, query.Error
}

func (r StatusRepository) FindByName(name string) (*StatusModel, error) {
	var status StatusModel
	query := r.DB.Where("name = ?", name).First(&status)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &status, query.Error
}

func (r StatusRepository) Create(status *StatusModel) error {
	query := r.DB.Create(&status)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r StatusRepository) Update(status_id uuid.UUID, status *StatusModel) error {
	query := r.DB.Model(StatusModel{}).Where("id", status_id).Updates(&status)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r StatusRepository) Destroy(status_id uuid.UUID) error {
	query := r.DB.Model(StatusModel{}).Delete(StatusModel{}, status_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
