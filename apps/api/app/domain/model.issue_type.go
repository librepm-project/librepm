package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueTypeModel struct {
	ID   uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name string
}

func (issue_type IssueTypeModel) TableName() string {
	return "issue_type"
}

func (issue_type *IssueTypeModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue_type.ID = uuid.New()
	return
}
