package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssueTypeStatusModel struct {
	ID                 uuid.UUID `gorm:"type:char(36);primary_key;"`
	IssueStatusID      uuid.UUID `gorm:"type:char(36) references issue_status;not null;"`
	IssueTypeID        uuid.UUID `gorm:"type:char(36) references issue_type;not null;"`
	ProjectIssueTypeID uuid.UUID `gorm:"type:char(36) references project_issue_type;not null;"`
	ProjectID          uuid.UUID `gorm:"type:char(36) references project;not null;"`

	IssueStatus      IssueTypeModel
	IssueType        IssueTypeModel
	ProjectIssueType ProjectIssueTypeModel
	Project          ProjectModel
}

func (project_issue_type_status ProjectIssueTypeStatusModel) TableName() string {
	return "project_issue_type_status"
}

func (project_issue_type_status *ProjectIssueTypeStatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_issue_type_status.ID = uuid.New()
	return
}
