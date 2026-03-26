package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttachmentModel struct {
	ID         uuid.UUID      `gorm:"type:char(36);primary_key;"`
	EntityType string         `gorm:"type:varchar(50);not null;"`
	EntityID   uuid.UUID      `gorm:"type:char(36);not null;"`
	Filename   string         `gorm:"type:varchar(255);not null;"`
	MimeType   string         `gorm:"type:varchar(100);not null;"`
	Size       int64          `gorm:"not null;"`
	StorePath  string         `gorm:"type:varchar(500);not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (a *AttachmentModel) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
