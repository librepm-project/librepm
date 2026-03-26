package domain

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type AttachmentServiceInterface interface {
	AllByEntity(entityType string, entityID uuid.UUID) (*[]AttachmentModel, error)
	FindByID(id uuid.UUID) (*AttachmentModel, error)
	Create(attachment *AttachmentModel, file multipart.File) error
	Destroy(id uuid.UUID) error
}

type AttachmentService struct {
	AttachmentRepository AttachmentRepositoryInterface
}

func (s AttachmentService) AllByEntity(entityType string, entityID uuid.UUID) (*[]AttachmentModel, error) {
	return s.AttachmentRepository.AllByEntity(entityType, entityID)
}

func (s AttachmentService) FindByID(id uuid.UUID) (*AttachmentModel, error) {
	return s.AttachmentRepository.FindByID(id)
}

func (s AttachmentService) Create(attachment *AttachmentModel, file multipart.File) error {
	dir := filepath.Join("uploads", attachment.EntityType, attachment.EntityID.String())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create upload directory: %w", err)
	}

	tmpID := uuid.New()
	storePath := filepath.Join(dir, tmpID.String()+"_"+attachment.Filename)

	dst, err := os.Create(storePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	attachment.StorePath = storePath
	return s.AttachmentRepository.Create(attachment)
}

func (s AttachmentService) Destroy(id uuid.UUID) error {
	attachment, err := s.AttachmentRepository.FindByID(id)
	if err != nil {
		return err
	}

	if err := s.AttachmentRepository.Destroy(id); err != nil {
		return err
	}

	os.Remove(attachment.StorePath)
	return nil
}
