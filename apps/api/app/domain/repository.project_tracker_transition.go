package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerTransitionRepositoryInterface interface {
	All() (*[]ProjectTrackerTransitionModel, error)
	FindByID(project_tracker_transition_id uuid.UUID) (*ProjectTrackerTransitionModel, error)
	Create(project_tracker_transition *ProjectTrackerTransitionModel) error
	Update(project_tracker_transition_id uuid.UUID, project_tracker_transition *ProjectTrackerTransitionModel) error
	Destroy(project_tracker_transition_id uuid.UUID) error
}

type ProjectTrackerTransitionRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerTransitionRepository) All() (*[]ProjectTrackerTransitionModel, error) {
	var project_tracker_transitions []ProjectTrackerTransitionModel
	var err error
	query := r.DB.Select("project_tracker_transition.*")

	if err := query.Find(&project_tracker_transitions).Error; err != nil {
		fmt.Println(err)
	}
	return &project_tracker_transitions, err
}

func (r ProjectTrackerTransitionRepository) FindByID(project_tracker_transition_id uuid.UUID) (*ProjectTrackerTransitionModel, error) {
	var project_tracker_transition ProjectTrackerTransitionModel
	query := r.DB.Model(ProjectTrackerTransitionModel{ID: project_tracker_transition_id}).Scan(&project_tracker_transition)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &project_tracker_transition, query.Error
}

func (r ProjectTrackerTransitionRepository) Create(project_tracker_transition *ProjectTrackerTransitionModel) error {
	query := r.DB.Create(&project_tracker_transition)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectTrackerTransitionRepository) Update(project_tracker_transition_id uuid.UUID, project_tracker_transition *ProjectTrackerTransitionModel) error {
	query := r.DB.Model(ProjectTrackerTransitionModel{}).Where("id", project_tracker_transition_id).Updates(&project_tracker_transition)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectTrackerTransitionRepository) Destroy(project_tracker_transition_id uuid.UUID) error {
	query := r.DB.Model(ProjectTrackerTransitionModel{}).Delete(ProjectTrackerTransitionModel{}, project_tracker_transition_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
