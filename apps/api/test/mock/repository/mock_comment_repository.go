package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockCommentRepository struct {
	AllByEntityFunc func(entityType string, entityID uuid.UUID) (*[]domain.CommentModel, error)
	FindByIDFunc    func(id uuid.UUID) (*domain.CommentModel, error)
	CreateFunc      func(comment *domain.CommentModel) error
	UpdateFunc      func(id uuid.UUID, comment *domain.CommentModel) error
	DestroyFunc     func(id uuid.UUID) error
}

func (m *MockCommentRepository) AllByEntity(entityType string, entityID uuid.UUID) (*[]domain.CommentModel, error) {
	if m.AllByEntityFunc != nil {
		return m.AllByEntityFunc(entityType, entityID)
	}
	return &[]domain.CommentModel{}, nil
}

func (m *MockCommentRepository) FindByID(id uuid.UUID) (*domain.CommentModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.CommentModel{}, nil
}

func (m *MockCommentRepository) Create(comment *domain.CommentModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(comment)
	}
	return nil
}

func (m *MockCommentRepository) Update(id uuid.UUID, comment *domain.CommentModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, comment)
	}
	return nil
}

func (m *MockCommentRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
