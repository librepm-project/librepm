package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssueTypeTransitionModel struct {
	ID                 uuid.UUID `gorm:"type:char(36);primary_key;"`
	IssueTypeID        uuid.UUID `gorm:"type:char(36) references issue_type;not null;"`
	IssueTransitionID  uuid.UUID `gorm:"type:char(36) references issue_transition;not null;"`
	ProjectIssueTypeID uuid.UUID `gorm:"type:char(36) references project_issue_type;not null;"`
	ProjectID          uuid.UUID `gorm:"type:char(36) references project;not null;"`

	IssueType        IssueTypeModel
	IssueTransition  IssueTransitionModel
	ProjectIssueType ProjectIssueTypeModel
	Project          ProjectModel
}

func (project_issue_type_transition ProjectIssueTypeTransitionModel) TableName() string {
	return "project_issue_type_transition"
}

func (project_issue_type_transition *ProjectIssueTypeTransitionModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_issue_type_transition.ID = uuid.New()
	return
}
