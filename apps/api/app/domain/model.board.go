package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name        string    `gorm:"type:varchar(100);not null;unique;"`
	Description string    `gorm:"type:varchar(255);"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	BoardColumns []BoardColumnModel `gorm:"foreignKey:BoardID;"`
}

func (board BoardModel) TableName() string {
	return "board"
}

func (board *BoardModel) BeforeCreate(tx *gorm.DB) (err error) {
	if board.ID == uuid.Nil {
		board.ID = uuid.New()
	}
	return
}
