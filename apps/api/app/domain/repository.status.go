package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusRepositoryInterface interface {
	All() *[]StatusModel
	FindByID(status_id uuid.UUID) *StatusModel
	Create(status *StatusModel)
	Update(status_id uuid.UUID, status *StatusModel)
	Destroy(status_id uuid.UUID)
}

type StatusRepository struct {
	DB *gorm.DB
}

func (r StatusRepository) All() *[]StatusModel {
	var statuses []StatusModel
	query := r.DB.Select("status.*")

	if err := query.Find(&statuses).Error; err != nil {
		fmt.Println(err)
	}
	return &statuses
}

func (r StatusRepository) FindByID(status_id uuid.UUID) *StatusModel {
	var status StatusModel
	fmt.Println(status_id)
	err := r.DB.Model(StatusModel{ID: status_id}).Scan(&status)

	if err != nil {
		fmt.Println(err)
	}
	return &status
}

func (r StatusRepository) Create(status *StatusModel) {
	r.DB.Create(&status)
}

func (r StatusRepository) Update(status_id uuid.UUID, status *StatusModel) {
	r.DB.Model(StatusModel{}).Where("id", status_id).Updates(&status)
}

func (r StatusRepository) Destroy(status_id uuid.UUID) {
	r.DB.Model(StatusModel{}).Delete(StatusModel{}, status_id)
}
