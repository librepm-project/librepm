package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectReleaseIssueRepositoryInterface interface {
	All() (*[]ProjectReleaseIssueModel, error)
	FindByID(id uuid.UUID) (*ProjectReleaseIssueModel, error)
	FindByProjectReleaseID(projectReleaseID uuid.UUID) (*[]ProjectReleaseIssueModel, error)
	FindByIssueID(issueID uuid.UUID) (*ProjectReleaseIssueModel, error)
	Create(pri *ProjectReleaseIssueModel) error
	Destroy(id uuid.UUID) error
}

type ProjectReleaseIssueRepository struct {
	DB *gorm.DB
}

func (r ProjectReleaseIssueRepository) All() (*[]ProjectReleaseIssueModel, error) {
	var prIssues []ProjectReleaseIssueModel
	if err := r.DB.Preload("Issue").Find(&prIssues).Error; err != nil {
		slog.Error("failed to fetch project release issues", "error", err)
		return nil, err
	}
	return &prIssues, nil
}

func (r ProjectReleaseIssueRepository) FindByID(id uuid.UUID) (*ProjectReleaseIssueModel, error) {
	var prIssue ProjectReleaseIssueModel
	if err := r.DB.Preload("Issue").First(&prIssue, "id = ?", id).Error; err != nil {
		slog.Error("failed to fetch project release issue", "id", id, "error", err)
		return nil, err
	}
	return &prIssue, nil
}

func (r ProjectReleaseIssueRepository) FindByProjectReleaseID(projectReleaseID uuid.UUID) (*[]ProjectReleaseIssueModel, error) {
	var prIssues []ProjectReleaseIssueModel
	if err := r.DB.Preload("Issue").Where("project_release_id = ?", projectReleaseID).Find(&prIssues).Error; err != nil {
		slog.Error("failed to fetch project release issues", "project_release_id", projectReleaseID, "error", err)
		return nil, err
	}
	return &prIssues, nil
}

func (r ProjectReleaseIssueRepository) FindByIssueID(issueID uuid.UUID) (*ProjectReleaseIssueModel, error) {
	var prIssue ProjectReleaseIssueModel
	if err := r.DB.Where("issue_id = ?", issueID).First(&prIssue).Error; err != nil {
		return nil, err
	}
	return &prIssue, nil
}

func (r ProjectReleaseIssueRepository) Create(pri *ProjectReleaseIssueModel) error {
	if err := r.DB.Create(pri).Error; err != nil {
		slog.Error("failed to create project release issue", "error", err)
		return err
	}
	return nil
}

func (r ProjectReleaseIssueRepository) Destroy(id uuid.UUID) error {
	if err := r.DB.Delete(&ProjectReleaseIssueModel{}, "id = ?", id).Error; err != nil {
		slog.Error("failed to delete project release issue", "id", id, "error", err)
		return err
	}
	return nil
}
