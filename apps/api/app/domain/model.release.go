package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ReleaseStatusPlanned  = "Planned"
	ReleaseStatusReleased = "Released"
)

type ReleaseModel struct {
	ID         uuid.UUID  `gorm:"type:char(36);primary_key;"`
	Name       string     `gorm:"type:varchar(255);not null;uniqueIndex:name_unique;"`
	Status     string     `gorm:"type:varchar(50);not null;default:'Planned';"`
	ReleasedAt *time.Time `gorm:"type:datetime;nullable;"`
	CreatedAt  time.Time  `gorm:"type:datetime;autoCreateTime;"`
	UpdatedAt  time.Time  `gorm:"type:datetime;autoUpdateTime;"`

	ProjectReleases []ProjectReleaseModel `gorm:"foreignKey:ReleaseID;"`
}

func (release ReleaseModel) TableName() string {
	return "release"
}

func (release *ReleaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	release.ID = uuid.New()
	return
}
