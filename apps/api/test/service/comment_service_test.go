package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCommentService_AllByEntity_ReturnsComments(t *testing.T) {
	entityID := uuid.New()
	expected := &[]domain.CommentModel{
		{ID: uuid.New(), Content: "Test comment"},
	}
	repo := &mockrepo.MockCommentRepository{
		AllByEntityFunc: func(entityType string, id uuid.UUID) (*[]domain.CommentModel, error) {
			assert.Equal(t, "issue", entityType)
			assert.Equal(t, entityID, id)
			return expected, nil
		},
	}
	svc := domain.CommentService{CommentRepository: repo}

	result, err := svc.AllByEntity("issue", entityID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestCommentService_AllByEntity_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockCommentRepository{
		AllByEntityFunc: func(_ string, _ uuid.UUID) (*[]domain.CommentModel, error) {
			return nil, repoErr
		},
	}
	svc := domain.CommentService{CommentRepository: repo}

	_, err := svc.AllByEntity("issue", uuid.New())

	assert.ErrorIs(t, err, repoErr)
}

func TestCommentService_FindByID_ReturnsComment(t *testing.T) {
	commentID := uuid.New()
	expected := &domain.CommentModel{ID: commentID, Content: "Hello"}
	repo := &mockrepo.MockCommentRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.CommentModel, error) {
			assert.Equal(t, commentID, id)
			return expected, nil
		},
	}
	svc := domain.CommentService{CommentRepository: repo}

	result, err := svc.FindByID(commentID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestCommentService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.CommentModel
	repo := &mockrepo.MockCommentRepository{
		CreateFunc: func(comment *domain.CommentModel) error {
			saved = comment
			return nil
		},
	}
	svc := domain.CommentService{CommentRepository: repo}

	comment := &domain.CommentModel{Content: "New comment", EntityType: "issue", EntityID: uuid.New()}
	err := svc.Create(comment)

	assert.NoError(t, err)
	assert.Equal(t, comment, saved)
}

func TestCommentService_Update_DelegatesCorrectly(t *testing.T) {
	commentID := uuid.New()
	var updatedID uuid.UUID
	repo := &mockrepo.MockCommentRepository{
		UpdateFunc: func(id uuid.UUID, comment *domain.CommentModel) error {
			updatedID = id
			return nil
		},
	}
	svc := domain.CommentService{CommentRepository: repo}

	err := svc.Update(commentID, &domain.CommentModel{Content: "Updated"})

	assert.NoError(t, err)
	assert.Equal(t, commentID, updatedID)
}

func TestCommentService_Destroy_DelegatesCorrectly(t *testing.T) {
	commentID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockCommentRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.CommentService{CommentRepository: repo}

	err := svc.Destroy(commentID)

	assert.NoError(t, err)
	assert.Equal(t, commentID, destroyedID)
}
