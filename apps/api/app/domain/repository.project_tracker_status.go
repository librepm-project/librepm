package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerStatusRepositoryInterface interface {
	All() *[]ProjectTrackerStatusModel
	FindByID(project_tracker_status_id uuid.UUID) *ProjectTrackerStatusModel
	Create(project_tracker_status *ProjectTrackerStatusModel)
	Update(project_tracker_status_id uuid.UUID, project_tracker_status *ProjectTrackerStatusModel)
	Destroy(project_tracker_status_id uuid.UUID)
}

type ProjectTrackerStatusRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerStatusRepository) All() *[]ProjectTrackerStatusModel {
	var project_tracker_statuses []ProjectTrackerStatusModel
	query := r.DB.Select("project_tracker_status.*")

	if err := query.Find(&project_tracker_statuses).Error; err != nil {
		fmt.Println(err)
	}
	return &project_tracker_statuses
}

func (r ProjectTrackerStatusRepository) FindByID(project_tracker_status_id uuid.UUID) *ProjectTrackerStatusModel {
	var project_tracker_status ProjectTrackerStatusModel
	fmt.Println(project_tracker_status_id)
	err := r.DB.Model(ProjectTrackerStatusModel{ID: project_tracker_status_id}).Scan(&project_tracker_status)

	if err != nil {
		fmt.Println(err)
	}
	return &project_tracker_status
}

func (r ProjectTrackerStatusRepository) Create(project_tracker_status *ProjectTrackerStatusModel) {
	r.DB.Create(&project_tracker_status)
}

func (r ProjectTrackerStatusRepository) Update(project_tracker_status_id uuid.UUID, project_tracker_status *ProjectTrackerStatusModel) {
	r.DB.Model(ProjectTrackerStatusModel{}).Where("id", project_tracker_status_id).Updates(&project_tracker_status)
}

func (r ProjectTrackerStatusRepository) Destroy(project_tracker_status_id uuid.UUID) {
	r.DB.Model(ProjectTrackerStatusModel{}).Delete(ProjectTrackerStatusModel{}, project_tracker_status_id)
}
