package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusModel struct {
	ID   uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name string    `gorm:"type:varchar(100);not null;unique;"`

	BoardColumnStatuses []BoardColumnStatusModel `gorm:"foreignKey:StatusID"`
}

func (status StatusModel) TableName() string {
	return "status"
}

func (status *StatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	status.ID = uuid.New()
	return
}
