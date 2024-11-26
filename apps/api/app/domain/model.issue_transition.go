package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueTransitionModel struct {
	ID                  uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name                string
	SourceIssueStatusID uuid.UUID `gorm:"type:char(36)"`
	SourceIssueStatus   IssueStatusModel
	TargetIssueStatusID uuid.UUID `gorm:"type:char(36)"`
	TargetIssueStatus   IssueStatusModel
}

func (issue_transition IssueTransitionModel) TableName() string {
	return "issue_transition"
}

func (issue_transition *IssueTransitionModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue_transition.ID = uuid.New()
	return
}
