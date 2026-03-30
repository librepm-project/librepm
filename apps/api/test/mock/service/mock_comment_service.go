package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockCommentService struct {
	AllByEntityFunc func(entityType string, entityID uuid.UUID) (*[]domain.CommentModel, error)
	FindByIDFunc    func(id uuid.UUID) (*domain.CommentModel, error)
	CreateFunc      func(comment *domain.CommentModel) error
	UpdateFunc      func(id uuid.UUID, comment *domain.CommentModel) error
	DestroyFunc     func(id uuid.UUID) error
}

func (m *MockCommentService) AllByEntity(entityType string, entityID uuid.UUID) (*[]domain.CommentModel, error) {
	if m.AllByEntityFunc != nil {
		return m.AllByEntityFunc(entityType, entityID)
	}
	return &[]domain.CommentModel{}, nil
}

func (m *MockCommentService) FindByID(id uuid.UUID) (*domain.CommentModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.CommentModel{}, nil
}

func (m *MockCommentService) Create(comment *domain.CommentModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(comment)
	}
	return nil
}

func (m *MockCommentService) Update(id uuid.UUID, comment *domain.CommentModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, comment)
	}
	return nil
}

func (m *MockCommentService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
