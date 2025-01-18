package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransitionModel struct {
	ID             uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name           string    `gorm:"varchar(100);not null;unique;"`
	SourceStatusID uuid.UUID `gorm:"type:char(36) references status;not null;uniqueIndex:source_status_id_target_status_id;"`
	TargetStatusID uuid.UUID `gorm:"type:char(36) rererences status;not null;uniqueIndex:source_status_id_target_status_id;"`

	SourceStatus StatusModel
	TargetStatus StatusModel
}

func (transition TransitionModel) TableName() string {
	return "transition"
}

func (transition *TransitionModel) BeforeCreate(tx *gorm.DB) (err error) {
	transition.ID = uuid.New()
	return
}
