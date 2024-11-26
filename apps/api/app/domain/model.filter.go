package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterModel struct {
	ID               uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name             string
	Description      string
	Public           bool
	UserID           uuid.UUID `gorm:"type:char(36) references user"`
	User             UserModel
	FilterConditions []FilterConditionModel `gorm:"foreignKey:FilterID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (filter FilterModel) TableName() string {
	return "filter"
}

func (filter *FilterModel) BeforeCreate(tx *gorm.DB) (err error) {
	filter.ID = uuid.New()
	return
}
