package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name        string
	Description string
	Public      bool
	UserID      uuid.UUID
	User        UserModel
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (dashboard DashboardModel) TableName() string {
	return "dashboard"

}

func (dashboard *DashboardModel) BeforeCreate(tx *gorm.DB) (err error) {
	dashboard.ID = uuid.New()
	return
}
