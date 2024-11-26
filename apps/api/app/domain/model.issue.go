package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueModel struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID     uuid.UUID
	Project       ProjectModel
	IssueTypeID   uuid.UUID `gorm:"type:char(36) references issue_type"`
	IssueType     IssueTypeModel
	IssueStatusID uuid.UUID `gorm:"type:char(36) references issue_status"`
	IssueStatus   IssueStatusModel
	ParentID      uuid.UUID `gorm:"type:char(36) references issue"`
	Parent        *IssueModel
	Key           string
	Summary       string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

func (issue IssueModel) TableName() string {
	return "issue"
}

func (issue *IssueModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue.ID = uuid.New()
	return
}
