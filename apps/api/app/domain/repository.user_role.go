package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleRepositoryInterface interface {
	FindByUserID(userID uuid.UUID) ([]UserRoleModel, error)
	Create(m *UserRoleModel) error
	Delete(userID uuid.UUID, role string) error
}

type UserRoleRepository struct {
	DB *gorm.DB
}

func (r UserRoleRepository) FindByUserID(userID uuid.UUID) ([]UserRoleModel, error) {
	var roles []UserRoleModel
	err := r.DB.Where("user_id = ?", userID).Find(&roles).Error
	if err != nil {
		slog.Error("failed to fetch user roles", "error", err)
	}
	return roles, err
}

func (r UserRoleRepository) Create(m *UserRoleModel) error {
	err := r.DB.Create(m).Error
	if err != nil {
		slog.Error("failed to create user role", "error", err)
	}
	return err
}

func (r UserRoleRepository) Delete(userID uuid.UUID, role string) error {
	err := r.DB.Where("user_id = ? AND role = ?", userID, role).Delete(&UserRoleModel{}).Error
	if err != nil {
		slog.Error("failed to delete user role", "error", err)
	}
	return err
}
