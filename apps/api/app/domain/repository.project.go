package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	All() *[]ProjectModel
	FindByID(project_id uuid.UUID) *ProjectModel
	Create(project *ProjectModel)
	Update(project_id uuid.UUID, project *ProjectModel)
	Destroy(project_id uuid.UUID)
}

type ProjectRepository struct {
	DB *gorm.DB
}

func (r ProjectRepository) All() *[]ProjectModel {
	var projects []ProjectModel
	query := r.DB.Select("project.*")

	if err := query.Find(&projects).Error; err != nil {
		fmt.Println(err)
	}
	return &projects
}

func (r ProjectRepository) FindByID(project_id uuid.UUID) *ProjectModel {
	var project ProjectModel
	fmt.Println(project_id)
	err := r.DB.Model(ProjectModel{ID: project_id}).Scan(&project)

	if err != nil {
		fmt.Println(err)
	}
	return &project
}

func (r ProjectRepository) Create(project *ProjectModel) {
	r.DB.Create(&project)
}

func (r ProjectRepository) Update(project_id uuid.UUID, project *ProjectModel) {
	r.DB.Model(ProjectModel{}).Where("id", project_id).Updates(&project)
}

func (r ProjectRepository) Destroy(project_id uuid.UUID) {
	r.DB.Model(ProjectModel{}).Delete(ProjectModel{}, project_id)
}
