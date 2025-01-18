package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerTransitionModel struct {
	ID               uuid.UUID `gorm:"type:char(36);primary_key;"`
	TransitionID     uuid.UUID `gorm:"type:char(36) references transition;not null;uniqueIndex:transition_id_project_tracker_id;"`
	ProjectTrackerID uuid.UUID `gorm:"type:char(36) references project_tracker;not null;uniqueIndex:transition_id_project_tracker_id;"`

	Transition     TransitionModel
	ProjectTracker ProjectTrackerModel
}

func (project_tracker_transition ProjectTrackerTransitionModel) TableName() string {
	return "project_tracker_transition"
}

func (project_tracker_transition *ProjectTrackerTransitionModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_tracker_transition.ID = uuid.New()
	return
}
