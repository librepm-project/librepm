package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	AllByEntity(entityType string, entityID uuid.UUID) (*[]CommentModel, error)
	FindByID(id uuid.UUID) (*CommentModel, error)
	Create(comment *CommentModel) error
	Update(id uuid.UUID, comment *CommentModel) error
	Destroy(id uuid.UUID) error
}

type CommentRepository struct {
	DB *gorm.DB
}

func (r CommentRepository) AllByEntity(entityType string, entityID uuid.UUID) (*[]CommentModel, error) {
	var comments []CommentModel
	err := r.DB.Preload("User").
		Where("entity_type = ? AND entity_id = ?", entityType, entityID).
		Order("created_at ASC").
		Find(&comments).Error
	if err != nil {
		slog.Error("failed to fetch comments", "error", err)
	}
	return &comments, err
}

func (r CommentRepository) FindByID(id uuid.UUID) (*CommentModel, error) {
	var comment CommentModel
	err := r.DB.Preload("User").First(&comment, id).Error
	if err != nil {
		slog.Error("failed to find comment by id", "error", err)
	}
	return &comment, err
}

func (r CommentRepository) Create(comment *CommentModel) error {
	err := r.DB.Create(comment).Error
	if err != nil {
		slog.Error("failed to create comment", "error", err)
	}
	return err
}

func (r CommentRepository) Update(id uuid.UUID, comment *CommentModel) error {
	err := r.DB.Model(CommentModel{}).Where("id = ?", id).Updates(comment).Error
	if err != nil {
		slog.Error("failed to update comment", "error", err)
	}
	return err
}

func (r CommentRepository) Destroy(id uuid.UUID) error {
	err := r.DB.Delete(&CommentModel{}, id).Error
	if err != nil {
		slog.Error("failed to destroy comment", "error", err)
	}
	return err
}
