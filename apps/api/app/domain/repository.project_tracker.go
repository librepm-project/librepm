package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerRepositoryInterface interface {
	All() *[]ProjectTrackerModel
	FindByID(project_tracker_id uuid.UUID) *ProjectTrackerModel
	Create(project_tracker *ProjectTrackerModel)
	Update(project_tracker_id uuid.UUID, project_tracker *ProjectTrackerModel)
	Destroy(project_tracker_id uuid.UUID)
}

type ProjectTrackerRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerRepository) All() *[]ProjectTrackerModel {
	var project_trackers []ProjectTrackerModel
	query := r.DB.Select("project_tracker.*")

	if err := query.Find(&project_trackers).Error; err != nil {
		fmt.Println(err)
	}
	return &project_trackers
}

func (r ProjectTrackerRepository) FindByID(project_tracker_id uuid.UUID) *ProjectTrackerModel {
	var project_tracker ProjectTrackerModel
	fmt.Println(project_tracker_id)
	err := r.DB.Model(ProjectTrackerModel{ID: project_tracker_id}).Scan(&project_tracker)

	if err != nil {
		fmt.Println(err)
	}
	return &project_tracker
}

func (r ProjectTrackerRepository) Create(project_tracker *ProjectTrackerModel) {
	r.DB.Create(&project_tracker)
}

func (r ProjectTrackerRepository) Update(project_tracker_id uuid.UUID, project_tracker *ProjectTrackerModel) {
	r.DB.Model(ProjectTrackerModel{}).Where("id", project_tracker_id).Updates(&project_tracker)
}

func (r ProjectTrackerRepository) Destroy(project_tracker_id uuid.UUID) {
	r.DB.Model(ProjectTrackerModel{}).Delete(ProjectTrackerModel{}, project_tracker_id)
}
