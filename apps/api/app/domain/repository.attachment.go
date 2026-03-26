package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttachmentRepositoryInterface interface {
	AllByEntity(entityType string, entityID uuid.UUID) (*[]AttachmentModel, error)
	FindByID(id uuid.UUID) (*AttachmentModel, error)
	Create(attachment *AttachmentModel) error
	Destroy(id uuid.UUID) error
}

type AttachmentRepository struct {
	DB *gorm.DB
}

func (r AttachmentRepository) AllByEntity(entityType string, entityID uuid.UUID) (*[]AttachmentModel, error) {
	var attachments []AttachmentModel
	err := r.DB.Where("entity_type = ? AND entity_id = ?", entityType, entityID).Find(&attachments).Error
	if err != nil {
		slog.Error("failed to fetch attachments", "error", err)
	}
	return &attachments, err
}

func (r AttachmentRepository) FindByID(id uuid.UUID) (*AttachmentModel, error) {
	var attachment AttachmentModel
	err := r.DB.First(&attachment, id).Error
	if err != nil {
		slog.Error("failed to find attachment by id", "error", err)
	}
	return &attachment, err
}

func (r AttachmentRepository) Create(attachment *AttachmentModel) error {
	err := r.DB.Create(attachment).Error
	if err != nil {
		slog.Error("failed to create attachment", "error", err)
	}
	return err
}

func (r AttachmentRepository) Destroy(id uuid.UUID) error {
	err := r.DB.Delete(&AttachmentModel{}, id).Error
	if err != nil {
		slog.Error("failed to destroy attachment", "error", err)
	}
	return err
}
