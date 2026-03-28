package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerStatusRepositoryInterface interface {
	All() (*[]ProjectTrackerStatusModel, error)
	FindByID(project_tracker_status_id uuid.UUID) (*ProjectTrackerStatusModel, error)
	FindStatusesByProjectID(project_id uuid.UUID) (*[]StatusModel, error)
	Create(project_tracker_status *ProjectTrackerStatusModel) error
	Update(project_tracker_status_id uuid.UUID, project_tracker_status *ProjectTrackerStatusModel) error
	Destroy(project_tracker_status_id uuid.UUID) error
}

type ProjectTrackerStatusRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerStatusRepository) All() (*[]ProjectTrackerStatusModel, error) {
	var project_tracker_statuses []ProjectTrackerStatusModel
	var err error
	query := r.DB.Select("project_tracker_status.*")

	if err = query.Find(&project_tracker_statuses).Error; err != nil {
		slog.Error("failed to fetch project tracker statuses", "error", err)
	}
	return &project_tracker_statuses, err
}

func (r ProjectTrackerStatusRepository) FindStatusesByProjectID(project_id uuid.UUID) (*[]StatusModel, error) {
	var statuses []StatusModel
	err := r.DB.
		Distinct("status.*").
		Joins("JOIN project_tracker_status pts ON pts.status_id = status.id").
		Joins("JOIN project_tracker pt ON pt.id = pts.project_tracker_id").
		Where("pt.project_id = ?", project_id).
		Find(&statuses).Error
	if err != nil {
		slog.Error("failed to find statuses by project id", "error", err)
	}
	return &statuses, err
}

func (r ProjectTrackerStatusRepository) FindByID(project_tracker_status_id uuid.UUID) (*ProjectTrackerStatusModel, error) {
	var project_tracker_status ProjectTrackerStatusModel
	query := r.DB.First(&project_tracker_status, project_tracker_status_id)

	if query.Error != nil {
		slog.Error("failed to find project tracker status by id", "error", query.Error)
	}
	return &project_tracker_status, query.Error
}

func (r ProjectTrackerStatusRepository) Create(project_tracker_status *ProjectTrackerStatusModel) error {
	query := r.DB.Create(&project_tracker_status)
	if query.Error != nil {
		slog.Error("failed to create project tracker status", "error", query.Error)
	}
	return query.Error
}

func (r ProjectTrackerStatusRepository) Update(project_tracker_status_id uuid.UUID, project_tracker_status *ProjectTrackerStatusModel) error {
	query := r.DB.Model(ProjectTrackerStatusModel{}).Where("id", project_tracker_status_id).Updates(&project_tracker_status)
	if query.Error != nil {
		slog.Error("failed to update project tracker status", "error", query.Error)
	}
	return query.Error
}

func (r ProjectTrackerStatusRepository) Destroy(project_tracker_status_id uuid.UUID) error {
	query := r.DB.Model(ProjectTrackerStatusModel{}).Delete(ProjectTrackerStatusModel{}, project_tracker_status_id)
	if query.Error != nil {
		slog.Error("failed to destroy project tracker status", "error", query.Error)
	}
	return query.Error
}
