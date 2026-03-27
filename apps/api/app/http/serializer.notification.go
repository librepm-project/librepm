package http

import (
	"apps/api/app/domain"
	"time"

	"github.com/google/uuid"
)

type NotificationResponse struct {
	ID         uuid.UUID  `json:"id"`
	Type       string     `json:"type"`
	EntityType *string    `json:"entityType,omitempty"`
	EntityID   *uuid.UUID `json:"entityId,omitempty"`
	Read       bool       `json:"read"`
	CreatedAt  time.Time  `json:"createdAt"`
}

type NotificationSerializer struct{}

func (s NotificationSerializer) ModelToResponse(n domain.NotificationModel) NotificationResponse {
	return NotificationResponse{
		ID:         n.ID,
		Type:       n.Type,
		EntityType: n.EntityType,
		EntityID:   n.EntityID,
		Read:       n.Read,
		CreatedAt:  n.CreatedAt,
	}
}

func (s NotificationSerializer) ModelToResponseMany(notifications []domain.NotificationModel) []NotificationResponse {
	result := []NotificationResponse{}
	for _, n := range notifications {
		result = append(result, s.ModelToResponse(n))
	}
	return result
}
