package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectTrackerRepositoryInterface interface {
	All() (*[]ProjectTrackerModel, error)
	FindByID(project_tracker_id uuid.UUID) (*ProjectTrackerModel, error)
	FindTrackersByProjectID(project_id uuid.UUID) (*[]TrackerModel, error)
	Create(project_tracker *ProjectTrackerModel) error
	Update(project_tracker_id uuid.UUID, project_tracker *ProjectTrackerModel) error
	Destroy(project_tracker_id uuid.UUID) error
}

type ProjectTrackerRepository struct {
	DB *gorm.DB
}

func (r ProjectTrackerRepository) All() (*[]ProjectTrackerModel, error) {
	var project_trackers []ProjectTrackerModel
	var err error
	query := r.DB.Select("project_tracker.*")

	if err := query.Find(&project_trackers).Error; err != nil {
		slog.Error("failed to fetch project trackers", "error", err)
	}
	return &project_trackers, err
}

func (r ProjectTrackerRepository) FindTrackersByProjectID(project_id uuid.UUID) (*[]TrackerModel, error) {
	var trackers []TrackerModel
	err := r.DB.
		Joins("JOIN project_tracker pt ON pt.tracker_id = tracker.id").
		Where("pt.project_id = ?", project_id).
		Find(&trackers).Error
	if err != nil {
		slog.Error("failed to find trackers by project id", "error", err)
	}
	return &trackers, err
}

func (r ProjectTrackerRepository) FindByID(project_tracker_id uuid.UUID) (*ProjectTrackerModel, error) {
	var project_tracker ProjectTrackerModel
	query := r.DB.First(&project_tracker, project_tracker_id)

	if query.Error != nil {
		slog.Error("failed to find project tracker by id", "error", query.Error)
	}
	return &project_tracker, query.Error
}

func (r ProjectTrackerRepository) Create(project_tracker *ProjectTrackerModel) error {
	query := r.DB.Create(&project_tracker)
	if query.Error != nil {
		slog.Error("failed to create project tracker", "error", query.Error)
	}
	return query.Error
}

func (r ProjectTrackerRepository) Update(project_tracker_id uuid.UUID, project_tracker *ProjectTrackerModel) error {
	query := r.DB.Model(ProjectTrackerModel{}).Where("id", project_tracker_id).Updates(&project_tracker)
	if query.Error != nil {
		slog.Error("failed to update project tracker", "error", query.Error)
	}
	return query.Error
}

func (r ProjectTrackerRepository) Destroy(project_tracker_id uuid.UUID) error {
	query := r.DB.Model(ProjectTrackerModel{}).Delete(ProjectTrackerModel{}, project_tracker_id)
	if query.Error != nil {
		slog.Error("failed to destroy project tracker", "error", query.Error)
	}
	return query.Error
}
