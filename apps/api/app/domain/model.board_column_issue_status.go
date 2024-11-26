package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnIssueStatusModel struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key;"`
	BoardColumnID uuid.UUID `gorm:"type:char(36) references board_column;not null;"`
	IssueStatusID uuid.UUID `gorm:"type:char(36) references issue_status;not null;"`

	BoardColumn BoardColumnModel
	IssueStatus IssueStatusModel
}

func (board_column_issue_status BoardColumnIssueStatusModel) TableName() string {
	return "board_column_issue_status"
}

func (board_column_issue_status *BoardColumnIssueStatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	board_column_issue_status.ID = uuid.New()
	return
}
