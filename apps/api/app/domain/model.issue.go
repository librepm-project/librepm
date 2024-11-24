package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID   uuid.UUID
	Key         string
	Type        string
	Summary     string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

func (issue IssueModel) TableName() string {
	return "issue"
}

func (issue *IssueModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue.ID = uuid.New()
	return
}
