package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssueTypeModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID   uuid.UUID `gorm:"type:char(36) references project;not null"`
	IssueTypeID uuid.UUID `gorm:"type:char(36) references issue_type;not null"`

	IssueType                   IssueTypeModel
	Project                     ProjectModel
	ProjectIssueTypeStatuses    []ProjectIssueTypeStatusModel     `gorm:"foreignKey:ProjectIssueTypeID;"`
	ProjectIssueTypeTransitions []ProjectIssueTypeTransitionModel `gorm:"foreignKey:ProjectIssueTypeID;"`
}

func (project_issue_type ProjectIssueTypeModel) TableName() string {
	return "project_issue_type"
}

func (project_issue_type *ProjectIssueTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_issue_type.ID = uuid.New()
	return
}
