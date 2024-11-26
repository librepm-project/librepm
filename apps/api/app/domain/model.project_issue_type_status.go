package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssueTypeStatusModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	StatusID  uuid.UUID
	Status    IssueTypeModel
	TypeID    uuid.UUID
	Type      IssueTypeModel
	ProjectID uuid.UUID
	Project   ProjectModel
}

func (project_issue_type_status ProjectIssueTypeStatusModel) TableName() string {
	return "project_issue_type_status"
}

func (project_issue_type_status *ProjectIssueTypeStatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_issue_type_status.ID = uuid.New()
	return
}
