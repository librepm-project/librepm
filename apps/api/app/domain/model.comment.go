package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentModel struct {
	ID         uuid.UUID      `gorm:"type:char(36);primary_key;"`
	EntityType string         `gorm:"type:varchar(50);not null;"`
	EntityID   uuid.UUID      `gorm:"type:char(36);not null;"`
	UserID     uuid.UUID      `gorm:"type:char(36);not null;"`
	Content    string         `gorm:"type:text;not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	User UserModel `gorm:"foreignKey:UserID"`
}

func (c *CommentModel) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
