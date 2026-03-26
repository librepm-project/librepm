package http

import (
	"time"

	"apps/api/app/domain"

	"github.com/google/uuid"
)

type AttachmentResponse struct {
	ID         uuid.UUID `json:"id"`
	Filename   string    `json:"filename"`
	MimeType   string    `json:"mimeType"`
	Size       int64     `json:"size"`
	EntityType string    `json:"entityType"`
	EntityID   uuid.UUID `json:"entityId"`
	CreatedAt  time.Time `json:"createdAt"`
}

type AttachmentSerializer struct{}

func (s AttachmentSerializer) ModelToResponse(a domain.AttachmentModel) AttachmentResponse {
	return AttachmentResponse{
		ID:         a.ID,
		Filename:   a.Filename,
		MimeType:   a.MimeType,
		Size:       a.Size,
		EntityType: a.EntityType,
		EntityID:   a.EntityID,
		CreatedAt:  a.CreatedAt,
	}
}

func (s AttachmentSerializer) ModelToResponseMany(attachments []domain.AttachmentModel) []AttachmentResponse {
	serialized := []AttachmentResponse{}
	for _, a := range attachments {
		serialized = append(serialized, s.ModelToResponse(a))
	}
	return serialized
}
