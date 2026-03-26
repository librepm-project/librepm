package http

import (
	"time"

	"apps/api/app/domain"

	"github.com/google/uuid"
)

type IssueAuditLogResponse struct {
	ID        uuid.UUID    `json:"id"`
	EventType string       `json:"eventType"`
	Field     string       `json:"field"`
	OldValue  string       `json:"oldValue"`
	NewValue  string       `json:"newValue"`
	Meta      string       `json:"meta"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"createdAt"`
}

type IssueAuditLogSerializer struct{}

func (s IssueAuditLogSerializer) ModelToResponse(l domain.IssueAuditLogModel) IssueAuditLogResponse {
	return IssueAuditLogResponse{
		ID:        l.ID,
		EventType: l.EventType,
		Field:     l.Field,
		OldValue:  l.OldValue,
		NewValue:  l.NewValue,
		Meta:      l.Meta,
		User:      UserSerializer{}.ModelToResponse(l.User),
		CreatedAt: l.CreatedAt,
	}
}

func (s IssueAuditLogSerializer) ModelToResponseMany(logs []domain.IssueAuditLogModel) []IssueAuditLogResponse {
	serialized := []IssueAuditLogResponse{}
	for _, l := range logs {
		serialized = append(serialized, s.ModelToResponse(l))
	}
	return serialized
}
