package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectReleaseIssueModel struct {
	ID                uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectReleaseID  uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:pr_issue_id;references:ProjectRelease(ID);"`
	IssueID           uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:pr_issue_id;references:Issue(ID);"`

	ProjectRelease ProjectReleaseModel
	Issue          IssueModel
}

func (pri ProjectReleaseIssueModel) TableName() string {
	return "project_release_issue"
}

func (pri *ProjectReleaseIssueModel) BeforeCreate(tx *gorm.DB) (err error) {
	pri.ID = uuid.New()
	return
}
