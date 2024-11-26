package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueStatusModel struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name        string    `gorm:"type:varchar(100);not null;"`
	IssueTypeID uuid.UUID `gorm:"type:char(36) references issue_type;not null;"`

	BoardColumnIssueStatuses []BoardColumnIssueStatusModel `gorm:"foreignKey:IssueStatusID"`
}

func (issue_status IssueStatusModel) TableName() string {
	return "issue_status"
}

func (issue_status *IssueStatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue_status.ID = uuid.New()
	return
}
