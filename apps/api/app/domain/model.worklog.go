package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueWorklogModel struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key;"`
	IssueID   uuid.UUID      `gorm:"type:char(36);not null;"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;"`
	Seconds   int            `gorm:"not null;"`
	Comment   string         `gorm:"type:varchar(255)"`
	LoggedAt  time.Time      `gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Issue IssueModel
	User  UserModel
}

func (w IssueWorklogModel) TableName() string {
	return "issue_worklog"
}

func (w *IssueWorklogModel) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = uuid.New()
	return
}
