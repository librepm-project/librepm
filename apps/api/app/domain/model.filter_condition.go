package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterConditionModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	Condition string    `gorm:"type:varchar(100);not null;uniqueIndex:condition_filter_id;"`
	FilterID  uuid.UUID `gorm:"type:char(36) references filter;not null;uniqueIndex:condition_filter_id;"`

	Filter FilterModel
}

func (filter_condition FilterConditionModel) TableName() string {
	return "filter_condition"
}

func (filter_condition *FilterConditionModel) BeforeCreate(tx *gorm.DB) (err error) {
	filter_condition.ID = uuid.New()
	return
}
