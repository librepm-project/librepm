package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockAttachmentRepository struct {
	AllByEntityFunc func(entityType string, entityID uuid.UUID) (*[]domain.AttachmentModel, error)
	FindByIDFunc    func(id uuid.UUID) (*domain.AttachmentModel, error)
	CreateFunc      func(attachment *domain.AttachmentModel) error
	DestroyFunc     func(id uuid.UUID) error
}

func (m *MockAttachmentRepository) AllByEntity(entityType string, entityID uuid.UUID) (*[]domain.AttachmentModel, error) {
	if m.AllByEntityFunc != nil {
		return m.AllByEntityFunc(entityType, entityID)
	}
	return &[]domain.AttachmentModel{}, nil
}

func (m *MockAttachmentRepository) FindByID(id uuid.UUID) (*domain.AttachmentModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.AttachmentModel{}, nil
}

func (m *MockAttachmentRepository) Create(attachment *domain.AttachmentModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(attachment)
	}
	return nil
}

func (m *MockAttachmentRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
