package domain

import (
	"log/slog"

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
	IncrementLastIssueKeyNumber(project_id uuid.UUID) error
}

type ProjectRepository struct {
	DB *gorm.DB
}

func (r ProjectRepository) All() (*[]ProjectModel, error) {
	var projects []ProjectModel
	var err error
	query := r.DB.Select("project.*")

	if err = query.Find(&projects).Error; err != nil {
		slog.Error("failed to fetch projects", "error", err)
	}
	return &projects, err
}

func (r ProjectRepository) FindByID(project_id uuid.UUID) (*ProjectModel, error) {
	var project ProjectModel
	query := r.DB.First(&project, project_id)

	if query.Error != nil {
		slog.Error("failed to find project by id", "error", query.Error)
	}
	return &project, query.Error
}

func (r ProjectRepository) FindByName(name string) (*ProjectModel, error) {
	var project ProjectModel
	query := r.DB.Where("name = ?", name).First(&project)

	if query.Error != nil {
		slog.Error("failed to find project by name", "error", query.Error)
	}
	return &project, query.Error
}

func (r ProjectRepository) Create(project *ProjectModel) error {
	query := r.DB.Create(&project)
	if query.Error != nil {
		slog.Error("failed to create project", "error", query.Error)
	}
	return query.Error
}

func (r ProjectRepository) Update(project_id uuid.UUID, project *ProjectModel) error {
	query := r.DB.Model(ProjectModel{}).Where("id", project_id).Updates(map[string]interface{}{
		"name":               project.Name,
		"code_name":          project.CodeName,
		"default_status_id":  project.DefaultStatusID,
		"default_tracker_id": project.DefaultTrackerID,
	})
	if query.Error != nil {
		slog.Error("failed to update project", "error", query.Error)
	}
	return query.Error
}

func (r ProjectRepository) Destroy(project_id uuid.UUID) error {
	query := r.DB.Model(ProjectModel{}).Delete(ProjectModel{}, project_id)
	if query.Error != nil {
		slog.Error("failed to destroy project", "error", query.Error)
	}
	return query.Error
}

func (r ProjectRepository) IncrementLastIssueKeyNumber(project_id uuid.UUID) error {
	query := r.DB.Model(&ProjectModel{}).Where("id = ?", project_id).Update("last_issue_key_number", gorm.Expr("last_issue_key_number + 1"))

	if query.Error != nil {
		slog.Error("failed to increment last issue key number", "error", query.Error)
		return query.Error
	}
	return nil
}
