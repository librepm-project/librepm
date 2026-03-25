package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectUserRepositoryInterface interface {
	All() (*[]ProjectUserModel, error)
	FindByID(project_user_id uuid.UUID) (*ProjectUserModel, error)
	Create(project_user *ProjectUserModel) error
	Update(project_user_id uuid.UUID, project_user *ProjectUserModel) error
	Destroy(project_user_id uuid.UUID) error
}

type ProjectUserRepository struct {
	DB *gorm.DB
}

func (r ProjectUserRepository) All() (*[]ProjectUserModel, error) {
	var project_users []ProjectUserModel
	var err error
	query := r.DB.Select("project_user.*")

	if err := query.Find(&project_users).Error; err != nil {
		slog.Error("failed to fetch project users", "error", err)
	}
	return &project_users, err
}

func (r ProjectUserRepository) FindByID(project_user_id uuid.UUID) (*ProjectUserModel, error) {
	var project_user ProjectUserModel
	query := r.DB.First(&project_user, project_user_id)

	if query.Error != nil {
		slog.Error("failed to find project user by id", "error", query.Error)
	}
	return &project_user, query.Error
}

func (r ProjectUserRepository) Create(project_user *ProjectUserModel) error {
	query := r.DB.Create(&project_user)
	if query.Error != nil {
		slog.Error("failed to create project user", "error", query.Error)
	}
	return query.Error
}

func (r ProjectUserRepository) Update(project_user_id uuid.UUID, project_user *ProjectUserModel) error {
	query := r.DB.Model(ProjectUserModel{}).Where("id", project_user_id).Updates(&project_user)
	if query.Error != nil {
		slog.Error("failed to update project user", "error", query.Error)
	}
	return query.Error
}

func (r ProjectUserRepository) Destroy(project_user_id uuid.UUID) error {
	query := r.DB.Model(ProjectUserModel{}).Delete(ProjectUserModel{}, project_user_id)
	if query.Error != nil {
		slog.Error("failed to destroy project user", "error", query.Error)
	}
	return query.Error
}
