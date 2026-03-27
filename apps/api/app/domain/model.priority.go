package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PriorityModel struct {
	ID    uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name  string    `gorm:"type:varchar(100);not null;unique;"`
	Color string    `gorm:"type:char(7);nullable;"`
}

func (priority PriorityModel) TableName() string {
	return "priority"
}

func (priority *PriorityModel) BeforeCreate(tx *gorm.DB) (err error) {
	priority.ID = uuid.New()
	return
}
