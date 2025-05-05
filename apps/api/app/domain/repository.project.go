package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	All() (*[]ProjectModel, error)
	FindByID(project_id uuid.UUID) (*ProjectModel, error)
	FindByName(name string) (*ProjectModel, error)
	Create(project *ProjectModel) error
	Update(project_id uuid.UUID, project *ProjectModel) error
	Destroy(project_id uuid.UUID) error
}

type ProjectRepository struct {
	DB *gorm.DB
}

func (r ProjectRepository) All() (*[]ProjectModel, error) {
	var projects []ProjectModel
	var err error
	query := r.DB.Select("project.*")

	if err := query.Find(&projects).Error; err != nil {
		fmt.Println(err)
	}
	return &projects, err
}

func (r ProjectRepository) FindByID(project_id uuid.UUID) (*ProjectModel, error) {
	var project ProjectModel
	query := r.DB.Model(ProjectModel{ID: project_id}).Scan(&project)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &project, query.Error
}

func (r ProjectRepository) FindByName(name string) (*ProjectModel, error) {
	var project ProjectModel
	query := r.DB.Where("name = ?", name).First(&project)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &project, query.Error
}

func (r ProjectRepository) Create(project *ProjectModel) error {
	query := r.DB.Create(&project)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectRepository) Update(project_id uuid.UUID, project *ProjectModel) error {
	query := r.DB.Model(ProjectModel{}).Where("id", project_id).Updates(&project)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r ProjectRepository) Destroy(project_id uuid.UUID) error {
	query := r.DB.Model(ProjectModel{}).Delete(ProjectModel{}, project_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
