package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectReleaseModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	ReleaseID uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:release_project_id;references:Release(ID);"`
	ProjectID uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:release_project_id;references:Project(ID);"`

	Release               ReleaseModel
	Project               ProjectModel
	ProjectReleaseIssues  []ProjectReleaseIssueModel `gorm:"foreignKey:ProjectReleaseID;"`
}

func (pr ProjectReleaseModel) TableName() string {
	return "project_release"
}

func (pr *ProjectReleaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	pr.ID = uuid.New()
	return
}
