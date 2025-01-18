package domain

import (
	"fmt"

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
		fmt.Println(err)
	}
	return &project_users, err
}

func (r ProjectUserRepository) FindByID(project_user_id uuid.UUID) (*ProjectUserModel, error) {
	var project_user ProjectUserModel
	query := r.DB.Model(ProjectUserModel{ID: project_user_id}).Scan(&project_user)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &project_user, query.Error
}

func (r ProjectUserRepository) Create(project_user *ProjectUserModel) error {
	query := r.DB.Create(&project_user)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectUserRepository) Update(project_user_id uuid.UUID, project_user *ProjectUserModel) error {
	query := r.DB.Model(ProjectUserModel{}).Where("id", project_user_id).Updates(&project_user)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectUserRepository) Destroy(project_user_id uuid.UUID) error {
	query := r.DB.Model(ProjectUserModel{}).Delete(ProjectUserModel{}, project_user_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
