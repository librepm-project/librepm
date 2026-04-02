package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleModel struct {
	ID     uuid.UUID `gorm:"type:char(36);primary_key;"`
	UserID uuid.UUID `gorm:"type:char(36);not null;index;uniqueIndex:idx_user_role;"`
	Role   string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_user_role;"`
}

func (m *UserRoleModel) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New()
	return nil
}
