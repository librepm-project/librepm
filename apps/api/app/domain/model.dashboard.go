package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name        string    `gorm:"type:varchar(100);not null;"`
	Description string    `gorm:"type:varchar(255);"`
	Public      bool
	UserID      uuid.UUID `gorm:"type:char(36) references user;not null;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User UserModel
}

func (dashboard DashboardModel) TableName() string {
	return "dashboard"

}

func (dashboard *DashboardModel) BeforeCreate(tx *gorm.DB) (err error) {
	dashboard.ID = uuid.New()
	return
}
