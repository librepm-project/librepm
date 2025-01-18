package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerStatusModel struct {
	ID               uuid.UUID `gorm:"type:char(36);primary_key;"`
	StatusID         uuid.UUID `gorm:"type:char(36) references status;not null;uniqueIndex:project_tracker_id_status_id;"`
	ProjectTrackerID uuid.UUID `gorm:"type:char(36) references project_tracker;not null;uniqueIndex:project_tracker_id_status_id;"`

	Status         StatusModel
	ProjectTracker ProjectTrackerModel
}

func (project_tracker_status ProjectTrackerStatusModel) TableName() string {
	return "project_tracker_status"
}

func (project_tracker_status *ProjectTrackerStatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_tracker_status.ID = uuid.New()
	return
}
