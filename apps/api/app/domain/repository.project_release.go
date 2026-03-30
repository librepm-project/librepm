package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectReleaseRepositoryInterface interface {
	All() (*[]ProjectReleaseModel, error)
	FindByID(id uuid.UUID) (*ProjectReleaseModel, error)
	FindByReleaseID(releaseID uuid.UUID) (*[]ProjectReleaseModel, error)
	FindByProjectID(projectID uuid.UUID) (*[]ProjectReleaseModel, error)
	Create(projectRelease *ProjectReleaseModel) error
	Update(id uuid.UUID, projectRelease *ProjectReleaseModel) error
	Destroy(id uuid.UUID) error
}

type ProjectReleaseRepository struct {
	DB *gorm.DB
}

func (r ProjectReleaseRepository) All() (*[]ProjectReleaseModel, error) {
	var projectReleases []ProjectReleaseModel
	if err := r.DB.Preload("Release").Preload("Project").Find(&projectReleases).Error; err != nil {
		slog.Error("failed to fetch project releases", "error", err)
		return nil, err
	}
	return &projectReleases, nil
}

func (r ProjectReleaseRepository) FindByID(id uuid.UUID) (*ProjectReleaseModel, error) {
	var projectRelease ProjectReleaseModel
	if err := r.DB.Preload("Release").Preload("Project").Preload("ProjectReleaseIssues").First(&projectRelease, "id = ?", id).Error; err != nil {
		slog.Error("failed to fetch project release", "id", id, "error", err)
		return nil, err
	}
	return &projectRelease, nil
}

func (r ProjectReleaseRepository) FindByReleaseID(releaseID uuid.UUID) (*[]ProjectReleaseModel, error) {
	var projectReleases []ProjectReleaseModel
	if err := r.DB.Preload("Release").Preload("Project").Where("release_id = ?", releaseID).Find(&projectReleases).Error; err != nil {
		slog.Error("failed to fetch project releases by release id", "release_id", releaseID, "error", err)
		return nil, err
	}
	return &projectReleases, nil
}

func (r ProjectReleaseRepository) FindByProjectID(projectID uuid.UUID) (*[]ProjectReleaseModel, error) {
	var projectReleases []ProjectReleaseModel
	if err := r.DB.Preload("Release").Preload("Project").Where("project_id = ?", projectID).Find(&projectReleases).Error; err != nil {
		slog.Error("failed to fetch project releases by project id", "project_id", projectID, "error", err)
		return nil, err
	}
	return &projectReleases, nil
}

func (r ProjectReleaseRepository) Create(projectRelease *ProjectReleaseModel) error {
	if err := r.DB.Create(projectRelease).Error; err != nil {
		slog.Error("failed to create project release", "error", err)
		return err
	}
	return nil
}

func (r ProjectReleaseRepository) Update(id uuid.UUID, projectRelease *ProjectReleaseModel) error {
	if err := r.DB.Model(&ProjectReleaseModel{}).Where("id = ?", id).Updates(projectRelease).Error; err != nil {
		slog.Error("failed to update project release", "id", id, "error", err)
		return err
	}
	return nil
}

func (r ProjectReleaseRepository) Destroy(id uuid.UUID) error {
	if err := r.DB.Delete(&ProjectReleaseModel{}, "id = ?", id).Error; err != nil {
		slog.Error("failed to delete project release", "id", id, "error", err)
		return err
	}
	return nil
}
