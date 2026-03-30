package mocksvc

import (
	"mime/multipart"

	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockAttachmentService struct {
	AllByEntityFunc func(entityType string, entityID uuid.UUID) (*[]domain.AttachmentModel, error)
	FindByIDFunc    func(id uuid.UUID) (*domain.AttachmentModel, error)
	CreateFunc      func(attachment *domain.AttachmentModel, file multipart.File) error
	DestroyFunc     func(id uuid.UUID) error
}

func (m *MockAttachmentService) AllByEntity(entityType string, entityID uuid.UUID) (*[]domain.AttachmentModel, error) {
	if m.AllByEntityFunc != nil {
		return m.AllByEntityFunc(entityType, entityID)
	}
	return &[]domain.AttachmentModel{}, nil
}

func (m *MockAttachmentService) FindByID(id uuid.UUID) (*domain.AttachmentModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.AttachmentModel{}, nil
}

func (m *MockAttachmentService) Create(attachment *domain.AttachmentModel, file multipart.File) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(attachment, file)
	}
	return nil
}

func (m *MockAttachmentService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
