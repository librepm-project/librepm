package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardModel struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name         string
	Description  string
	BoardColumns []BoardColumnModel `gorm:"foreignKey:BoardID;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (board BoardModel) TableName() string {
	return "board"
}

func (board *BoardModel) BeforeCreate(tx *gorm.DB) (err error) {
	board.ID = uuid.New()
	return
}
