package http

import (
	"time"

	"apps/api/app/domain"

	"github.com/google/uuid"
)

type CommentRequest struct {
	Content string `json:"content"`
}

type CommentResponse struct {
	ID         uuid.UUID    `json:"id"`
	Content    string       `json:"content"`
	EntityType string       `json:"entityType"`
	EntityID   uuid.UUID    `json:"entityId"`
	User       UserResponse `json:"user"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}

type CommentSerializer struct{}

func (s CommentSerializer) ModelToResponse(c domain.CommentModel) CommentResponse {
	return CommentResponse{
		ID:         c.ID,
		Content:    c.Content,
		EntityType: c.EntityType,
		EntityID:   c.EntityID,
		User:       UserSerializer{}.ModelToResponse(c.User),
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func (s CommentSerializer) ModelToResponseMany(comments []domain.CommentModel) []CommentResponse {
	serialized := []CommentResponse{}
	for _, c := range comments {
		serialized = append(serialized, s.ModelToResponse(c))
	}
	return serialized
}
