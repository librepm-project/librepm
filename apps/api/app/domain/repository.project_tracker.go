package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerRepositoryInterface interface {
	All() (*[]ProjectTrackerModel, error)
	FindByID(project_tracker_id uuid.UUID) (*ProjectTrackerModel, error)
	Create(project_tracker *ProjectTrackerModel) error
	Update(project_tracker_id uuid.UUID, project_tracker *ProjectTrackerModel) error
	Destroy(project_tracker_id uuid.UUID) error
}

type ProjectTrackerRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerRepository) All() (*[]ProjectTrackerModel, error) {
	var project_trackers []ProjectTrackerModel
	var err error
	query := r.DB.Select("project_tracker.*")

	if err := query.Find(&project_trackers).Error; err != nil {
		fmt.Println(err)
	}
	return &project_trackers, err
}

func (r ProjectTrackerRepository) FindByID(project_tracker_id uuid.UUID) (*ProjectTrackerModel, error) {
	var project_tracker ProjectTrackerModel
	query := r.DB.Model(ProjectTrackerModel{ID: project_tracker_id}).Scan(&project_tracker)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &project_tracker, query.Error
}

func (r ProjectTrackerRepository) Create(project_tracker *ProjectTrackerModel) error {
	query := r.DB.Create(&project_tracker)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectTrackerRepository) Update(project_tracker_id uuid.UUID, project_tracker *ProjectTrackerModel) error {
	query := r.DB.Model(ProjectTrackerModel{}).Where("id", project_tracker_id).Updates(&project_tracker)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectTrackerRepository) Destroy(project_tracker_id uuid.UUID) error {
	query := r.DB.Model(ProjectTrackerModel{}).Delete(ProjectTrackerModel{}, project_tracker_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
