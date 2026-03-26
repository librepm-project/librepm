package domain

import (
	"github.com/google/uuid"
)

type CommentServiceInterface interface {
	AllByEntity(entityType string, entityID uuid.UUID) (*[]CommentModel, error)
	FindByID(id uuid.UUID) (*CommentModel, error)
	Create(comment *CommentModel) error
	Update(id uuid.UUID, comment *CommentModel) error
	Destroy(id uuid.UUID) error
}

type CommentService struct {
	CommentRepository CommentRepositoryInterface
}

func (s CommentService) AllByEntity(entityType string, entityID uuid.UUID) (*[]CommentModel, error) {
	return s.CommentRepository.AllByEntity(entityType, entityID)
}

func (s CommentService) FindByID(id uuid.UUID) (*CommentModel, error) {
	return s.CommentRepository.FindByID(id)
}

func (s CommentService) Create(comment *CommentModel) error {
	return s.CommentRepository.Create(comment)
}

func (s CommentService) Update(id uuid.UUID, comment *CommentModel) error {
	return s.CommentRepository.Update(id, comment)
}

func (s CommentService) Destroy(id uuid.UUID) error {
	return s.CommentRepository.Destroy(id)
}
