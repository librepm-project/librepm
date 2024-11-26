package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssueTypeTransitionModel struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key;"`
	TypeID       uuid.UUID
	Type         IssueTypeModel
	TransitionID uuid.UUID
	Transition   IssueTransitionModel
	ProjectID    uuid.UUID
	Project      ProjectModel
}

func (project_issue_type_transition ProjectIssueTypeTransitionModel) TableName() string {
	return "project_issue_type_transition"
}

func (project_issue_type_transition *ProjectIssueTypeTransitionModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_issue_type_transition.ID = uuid.New()
	return
}
