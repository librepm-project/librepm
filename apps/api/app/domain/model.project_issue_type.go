package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssueTypeModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID   uuid.UUID
	Project     ProjectModel
	IssueTypeID uuid.UUID
	IssueType   IssueTypeModel
}

func (project_issue_type ProjectIssueTypeModel) TableName() string {
	return "project_issue_type"
}

func (project_issue_type *ProjectIssueTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_issue_type.ID = uuid.New()
	return
}
