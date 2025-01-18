package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerStatusRepositoryInterface interface {
	All() (*[]ProjectTrackerStatusModel, error)
	FindByID(project_tracker_status_id uuid.UUID) (*ProjectTrackerStatusModel, error)
	Create(project_tracker_status *ProjectTrackerStatusModel) error
	Update(project_tracker_status_id uuid.UUID, project_tracker_status *ProjectTrackerStatusModel) error
	Destroy(project_tracker_status_id uuid.UUID) error
}

type ProjectTrackerStatusRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerStatusRepository) All() (*[]ProjectTrackerStatusModel, error) {
	var project_tracker_statuses []ProjectTrackerStatusModel
	var err error
	query := r.DB.Select("project_tracker_status.*")

	if err := query.Find(&project_tracker_statuses).Error; err != nil {
		fmt.Println(err)
	}
	return &project_tracker_statuses, err
}

func (r ProjectTrackerStatusRepository) FindByID(project_tracker_status_id uuid.UUID) (*ProjectTrackerStatusModel, error) {
	var project_tracker_status ProjectTrackerStatusModel
	query := r.DB.Model(ProjectTrackerStatusModel{ID: project_tracker_status_id}).Scan(&project_tracker_status)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &project_tracker_status, query.Error
}

func (r ProjectTrackerStatusRepository) Create(project_tracker_status *ProjectTrackerStatusModel) error {
	query := r.DB.Create(&project_tracker_status)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectTrackerStatusRepository) Update(project_tracker_status_id uuid.UUID, project_tracker_status *ProjectTrackerStatusModel) error {
	query := r.DB.Model(ProjectTrackerStatusModel{}).Where("id", project_tracker_status_id).Updates(&project_tracker_status)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectTrackerStatusRepository) Destroy(project_tracker_status_id uuid.UUID) error {
	query := r.DB.Model(ProjectTrackerStatusModel{}).Delete(ProjectTrackerStatusModel{}, project_tracker_status_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
