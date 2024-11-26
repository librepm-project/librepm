package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name        string    `gorm:"type:varchar(100);not null;"`
	Description string    `gorm:"type:varchar(255);"`
	Public      bool
	UserID      uuid.UUID `gorm:"type:char(36) references user;not null;"`
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
