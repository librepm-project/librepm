package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID uuid.UUID `gorm:"type:char(36) references project;not null;uniqueIndex:project_id_tracker_id;"`
	TrackerID uuid.UUID `gorm:"type:char(36) references tracker;not null;uniqueIndex:project_id_tracker_id;"`

	Tracker                   TrackerModel
	Project                   ProjectModel
	ProjectTrackerStatuses    []ProjectTrackerStatusModel     `gorm:"foreignKey:ProjectTrackerID;"`
	ProjectTrackerTransitions []ProjectTrackerTransitionModel `gorm:"foreignKey:ProjectTrackerID;"`
}

func (project_tracker ProjectTrackerModel) TableName() string {
	return "project_tracker"
}

func (project_tracker *ProjectTrackerModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_tracker.ID = uuid.New()
	return
}
