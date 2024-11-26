package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectUserModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	ProjectID uuid.UUID `gorm:"type:char(36);not null;"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;"`

	Project ProjectModel
	User    UserModel
}

func (project_user ProjectUserModel) TableName() string {
	return "project_user"
}

func (project_user *ProjectUserModel) BeforeCreate(tx *gorm.DB) (err error) {
	project_user.ID = uuid.New()
	return
}
