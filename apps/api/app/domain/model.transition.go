package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransitionModel struct {
	ID             uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name           string    `gorm:"varchar(100);not null;"`
	SourceStatusID uuid.UUID `gorm:"type:char(36) references status;not null;"`
	TargetStatusID uuid.UUID `gorm:"type:char(36) rererences status;not null;"`

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
