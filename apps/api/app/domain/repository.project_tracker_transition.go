package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerTransitionRepositoryInterface interface {
	All() *[]ProjectTrackerTransitionModel
	FindByID(project_tracker_transition_id uuid.UUID) *ProjectTrackerTransitionModel
	Create(project_tracker_transition *ProjectTrackerTransitionModel)
	Update(project_tracker_transition_id uuid.UUID, project_tracker_transition *ProjectTrackerTransitionModel)
	Destroy(project_tracker_transition_id uuid.UUID)
}

type ProjectTrackerTransitionRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerTransitionRepository) All() *[]ProjectTrackerTransitionModel {
	var project_tracker_transitions []ProjectTrackerTransitionModel
	query := r.DB.Select("project_tracker_transition.*")

	if err := query.Find(&project_tracker_transitions).Error; err != nil {
		fmt.Println(err)
	}
	return &project_tracker_transitions
}

func (r ProjectTrackerTransitionRepository) FindByID(project_tracker_transition_id uuid.UUID) *ProjectTrackerTransitionModel {
	var project_tracker_transition ProjectTrackerTransitionModel
	fmt.Println(project_tracker_transition_id)
	err := r.DB.Model(ProjectTrackerTransitionModel{ID: project_tracker_transition_id}).Scan(&project_tracker_transition)

	if err != nil {
		fmt.Println(err)
	}
	return &project_tracker_transition
}

func (r ProjectTrackerTransitionRepository) Create(project_tracker_transition *ProjectTrackerTransitionModel) {
	r.DB.Create(&project_tracker_transition)
}

func (r ProjectTrackerTransitionRepository) Update(project_tracker_transition_id uuid.UUID, project_tracker_transition *ProjectTrackerTransitionModel) {
	r.DB.Model(ProjectTrackerTransitionModel{}).Where("id", project_tracker_transition_id).Updates(&project_tracker_transition)
}

func (r ProjectTrackerTransitionRepository) Destroy(project_tracker_transition_id uuid.UUID) {
	r.DB.Model(ProjectTrackerTransitionModel{}).Delete(ProjectTrackerTransitionModel{}, project_tracker_transition_id)
}
