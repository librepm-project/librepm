package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnModel struct {
	ID                       uuid.UUID `gorm:"type:char(36);primary_key;"`
	BoardID                  uuid.UUID `gorm:"type:char(36) references board"`
	Board                    BoardModel
	Label                    string
	BoardColumnIssueStatuses []BoardColumnIssueStatusModel `gorm:"foreignKey:BoardColumnID"`
}

func (board_column BoardColumnModel) TableName() string {
	return "board_column"
}

func (board_column *BoardColumnModel) BeforeCreate(tx *gorm.DB) (err error) {
	board_column.ID = uuid.New()
	return
}
