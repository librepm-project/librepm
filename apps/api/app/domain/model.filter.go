package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex:name_user_id;"`
	Description string    `gorm:"type:varchar(255);"`
	Public      bool
	ColumnList  string    `gorm:"type:json;default:'[\"key\", \"tracker\", \"priority\", \"summary\", \"assignee\", \"status\"]'"`
	GroupBy     string    `gorm:"type:varchar(50);default:''"`
	UserID      uuid.UUID `gorm:"type:char(36) references user;not null;uniqueIndex:name_user_id;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User             UserModel
	FilterConditions []FilterConditionModel `gorm:"foreignKey:FilterID"`
}

func (filter FilterModel) TableName() string {
	return "filter"
}

func (filter *FilterModel) BeforeCreate(tx *gorm.DB) (err error) {
	filter.ID = uuid.New()
	return
}
