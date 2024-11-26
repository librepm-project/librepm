package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueModel struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID     uuid.UUID `gorm:"type:char(36);not null;"`
	IssueTypeID   uuid.UUID `gorm:"type:char(36) references issue_type;not null;"`
	IssueStatusID uuid.UUID `gorm:"type:char(36) references issue_status;not null;"`
	ParentID      uuid.UUID `gorm:"type:char(36) references issue;not null;"`
	Key           string    `gorm:"type:varchar(5);not null;"`
	Summary       string    `gorm:"type:varchar(100);not null;"`
	Description   string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`

	Project     ProjectModel
	IssueType   IssueTypeModel
	IssueStatus IssueStatusModel

	Parent *IssueModel
}

func (issue IssueModel) TableName() string {
	return "issue"
}

func (issue *IssueModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue.ID = uuid.New()
	return
}
