package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardWidgetModel struct {
	ID          uuid.UUID  `gorm:"type:char(36);primary_key;"`
	DashboardID uuid.UUID  `gorm:"type:char(36) references dashboard;not null;"`
	FilterID    *uuid.UUID `gorm:"type:char(36) references filter;"`
	Type        string     `gorm:"type:varchar(50);not null;"`
	Title       string     `gorm:"type:varchar(100);not null;"`
	CreatedAt   time.Time

	Dashboard DashboardModel
	Filter    *FilterModel
}

func (dashboard_widget DashboardWidgetModel) TableName() string {
	return "dashboard_widget"
}

func (dashboard_widget *DashboardWidgetModel) BeforeCreate(tx *gorm.DB) (err error) {
	dashboard_widget.ID = uuid.New()
	return
}
