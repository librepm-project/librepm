package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationModel struct {
	ID         uuid.UUID  `gorm:"type:char(36);primary_key;"`
	UserID     uuid.UUID  `gorm:"type:char(36);not null;index;"`
	Type       string     `gorm:"type:varchar(50);not null;"`
	EntityType *string    `gorm:"type:varchar(50);"`
	EntityID   *uuid.UUID `gorm:"type:char(36);"`
	Read       bool       `gorm:"default:false;not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (n *NotificationModel) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New()
	return
}
