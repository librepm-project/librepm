package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectModel struct {
	ID                          uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name                        string
	CodeName                    string
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	DeletedAt                   *time.Time                        `sql:"index"`
	Issues                      []IssueModel                      `gorm:"foreignKey:ProjectID"`
	ProjectUsers                []ProjectUserModel                `gorm:"foreignKey:ProjectID"`
	ProjectIssueTypeStatuses    []ProjectIssueTypeStatusModel     `gorm:"foreignKey:ProjectID"`
	ProjectIssueTypeTransitions []ProjectIssueTypeTransitionModel `gorm:"foreignKey:ProjectID"`
}

func (project ProjectModel) TableName() string {
	return "project"
}

func (project *ProjectModel) BeforeCreate(tx *gorm.DB) (err error) {
	project.ID = uuid.New()
	return
}
