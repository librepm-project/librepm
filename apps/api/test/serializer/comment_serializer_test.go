package serializer_test

import (
	"testing"
	"time"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCommentSerializer_ModelToResponse_MapsFields(t *testing.T) {
	commentID := uuid.New()
	entityID := uuid.New()
	userID := uuid.New()
	now := time.Now()

	model := domain.CommentModel{
		ID:         commentID,
		Content:    "This is a comment",
		EntityType: "issue",
		EntityID:   entityID,
		User:       domain.UserModel{ID: userID, Email: "user@example.com"},
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	resp := apphttp.CommentSerializer{}.ModelToResponse(model)

	assert.Equal(t, commentID, resp.ID)
	assert.Equal(t, "This is a comment", resp.Content)
	assert.Equal(t, "issue", resp.EntityType)
	assert.Equal(t, entityID, resp.EntityID)
	assert.Equal(t, userID, resp.User.ID)
}

func TestCommentSerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.CommentModel{
		{ID: uuid.New(), Content: "First comment"},
		{ID: uuid.New(), Content: "Second comment"},
	}

	resp := apphttp.CommentSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, "First comment", resp[0].Content)
}

func TestCommentSerializer_ModelToResponseMany_EmptySlice(t *testing.T) {
	resp := apphttp.CommentSerializer{}.ModelToResponseMany([]domain.CommentModel{})

	assert.Empty(t, resp)
}
