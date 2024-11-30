package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectUserRepositoryInterface interface {
	All() *[]ProjectUserModel
	FindByID(project_user_id uuid.UUID) *ProjectUserModel
	Create(project_user *ProjectUserModel)
	Update(project_user_id uuid.UUID, project_user *ProjectUserModel)
	Destroy(project_user_id uuid.UUID)
}

type ProjectUserRepository struct {
	DB *gorm.DB
}

func (r ProjectUserRepository) All() *[]ProjectUserModel {
	var project_users []ProjectUserModel
	query := r.DB.Select("project_user.*")

	if err := query.Find(&project_users).Error; err != nil {
		fmt.Println(err)
	}
	return &project_users
}

func (r ProjectUserRepository) FindByID(project_user_id uuid.UUID) *ProjectUserModel {
	var project_user ProjectUserModel
	fmt.Println(project_user_id)
	err := r.DB.Model(ProjectUserModel{ID: project_user_id}).Scan(&project_user)

	if err != nil {
		fmt.Println(err)
	}
	return &project_user
}

func (r ProjectUserRepository) Create(project_user *ProjectUserModel) {
	r.DB.Create(&project_user)
}

func (r ProjectUserRepository) Update(project_user_id uuid.UUID, project_user *ProjectUserModel) {
	r.DB.Model(ProjectUserModel{}).Where("id", project_user_id).Updates(&project_user)
}

func (r ProjectUserRepository) Destroy(project_user_id uuid.UUID) {
	r.DB.Model(ProjectUserModel{}).Delete(ProjectUserModel{}, project_user_id)
}
