package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrackerRepositoryInterface interface {
	All() (*[]TrackerModel, error)
	FindByID(tracker_id uuid.UUID) (*TrackerModel, error)
	FindByName(name string) (*TrackerModel, error)
	Create(tracker *TrackerModel) error
	Update(tracker_id uuid.UUID, tracker *TrackerModel) error
	Destroy(tracker_id uuid.UUID) error
}

type TrackerRepository struct {
	DB *gorm.DB
}

func (r TrackerRepository) All() (*[]TrackerModel, error) {
	var trackers []TrackerModel
	var err error
	query := r.DB.Select("tracker.*")

	if err := query.Find(&trackers).Error; err != nil {
		fmt.Println(err)
	}
	return &trackers, err
}

func (r TrackerRepository) FindByID(tracker_id uuid.UUID) (*TrackerModel, error) {
	var tracker TrackerModel
	query := r.DB.Model(TrackerModel{ID: tracker_id}).Scan(&tracker)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &tracker, query.Error
}

func (r TrackerRepository) FindByName(name string) (*TrackerModel, error) {
	var tracker TrackerModel
	query := r.DB.Where("name = ?", name).First(&tracker)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &tracker, query.Error
}

func (r TrackerRepository) Create(tracker *TrackerModel) error {
	query := r.DB.Create(&tracker)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r TrackerRepository) Update(tracker_id uuid.UUID, tracker *TrackerModel) error {
	query := r.DB.Model(TrackerModel{}).Where("id", tracker_id).Updates(&tracker)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r TrackerRepository) Destroy(tracker_id uuid.UUID) error {
	query := r.DB.Model(TrackerModel{}).Delete(TrackerModel{}, tracker_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
