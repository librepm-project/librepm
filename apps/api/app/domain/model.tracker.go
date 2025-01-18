package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrackerModel struct {
	ID   uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name string    `gorm:"type:varchar(100);not null;unique;"`
}

func (tracker TrackerModel) TableName() string {
	return "tracker"
}

func (tracker *TrackerModel) BeforeCreate(tx *gorm.DB) (err error) {
	tracker.ID = uuid.New()
	return
}
