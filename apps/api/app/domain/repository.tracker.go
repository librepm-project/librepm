package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrackerRepositoryInterface interface {
	All() *[]TrackerModel
	FindByID(tracker_id uuid.UUID) *TrackerModel
	Create(tracker *TrackerModel)
	Update(tracker_id uuid.UUID, tracker *TrackerModel)
	Destroy(tracker_id uuid.UUID)
}

type TrackerRepository struct {
	DB *gorm.DB
}

func (r TrackerRepository) All() *[]TrackerModel {
	var trackers []TrackerModel
	query := r.DB.Select("tracker.*")

	if err := query.Find(&trackers).Error; err != nil {
		fmt.Println(err)
	}
	return &trackers
}

func (r TrackerRepository) FindByID(tracker_id uuid.UUID) *TrackerModel {
	var tracker TrackerModel
	fmt.Println(tracker_id)
	err := r.DB.Model(TrackerModel{ID: tracker_id}).Scan(&tracker)

	if err != nil {
		fmt.Println(err)
	}
	return &tracker
}

func (r TrackerRepository) Create(tracker *TrackerModel) {
	r.DB.Create(&tracker)
}

func (r TrackerRepository) Update(tracker_id uuid.UUID, tracker *TrackerModel) {
	r.DB.Model(TrackerModel{}).Where("id", tracker_id).Updates(&tracker)
}

func (r TrackerRepository) Destroy(tracker_id uuid.UUID) {
	r.DB.Model(TrackerModel{}).Delete(TrackerModel{}, tracker_id)
}
