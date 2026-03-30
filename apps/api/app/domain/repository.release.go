package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReleaseRepositoryInterface interface {
	All() (*[]ReleaseModel, error)
	FindByID(id uuid.UUID) (*ReleaseModel, error)
	Create(release *ReleaseModel) error
	Update(id uuid.UUID, release *ReleaseModel) error
	Destroy(id uuid.UUID) error
}

type ReleaseRepository struct {
	DB *gorm.DB
}

func (r ReleaseRepository) All() (*[]ReleaseModel, error) {
	var releases []ReleaseModel
	if err := r.DB.Find(&releases).Error; err != nil {
		slog.Error("failed to fetch releases", "error", err)
		return nil, err
	}
	return &releases, nil
}

func (r ReleaseRepository) FindByID(id uuid.UUID) (*ReleaseModel, error) {
	var release ReleaseModel
	if err := r.DB.Preload("ProjectReleases").First(&release, "id = ?", id).Error; err != nil {
		slog.Error("failed to fetch release", "id", id, "error", err)
		return nil, err
	}
	return &release, nil
}

func (r ReleaseRepository) Create(release *ReleaseModel) error {
	if err := r.DB.Create(release).Error; err != nil {
		slog.Error("failed to create release", "error", err)
		return err
	}
	return nil
}

func (r ReleaseRepository) Update(id uuid.UUID, release *ReleaseModel) error {
	if err := r.DB.Model(&ReleaseModel{}).Where("id = ?", id).Updates(release).Error; err != nil {
		slog.Error("failed to update release", "id", id, "error", err)
		return err
	}
	return nil
}

func (r ReleaseRepository) Destroy(id uuid.UUID) error {
	if err := r.DB.Delete(&ReleaseModel{}, "id = ?", id).Error; err != nil {
		slog.Error("failed to delete release", "id", id, "error", err)
		return err
	}
	return nil
}
