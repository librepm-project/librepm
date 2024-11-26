package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueTransitionModel struct {
	ID                  uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name                string    `gorm:"varchar(100);not null;"`
	SourceIssueStatusID uuid.UUID `gorm:"type:char(36) references issue_status;not null;"`
	TargetIssueStatusID uuid.UUID `gorm:"type:char(36) rererences issue_status;not null;"`

	SourceIssueStatus IssueStatusModel
	TargetIssueStatus IssueStatusModel
}

func (issue_transition IssueTransitionModel) TableName() string {
	return "issue_transition"
}

func (issue_transition *IssueTransitionModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue_transition.ID = uuid.New()
	return
}
