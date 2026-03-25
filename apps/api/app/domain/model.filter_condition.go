package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterConditionModel struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key;"`
	FilterID uuid.UUID `gorm:"type:char(36) references filter;not null;"`
	Field    string    `gorm:"type:varchar(50);not null;"`
	Op       string    `gorm:"type:varchar(20);not null;"`
	Value    string    `gorm:"type:varchar(255);not null;"`

	Filter FilterModel
}

func (filter_condition FilterConditionModel) TableName() string {
	return "filter_condition"
}

func (filter_condition *FilterConditionModel) BeforeCreate(tx *gorm.DB) (err error) {
	filter_condition.ID = uuid.New()
	return
}
