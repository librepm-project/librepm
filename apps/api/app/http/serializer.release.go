package http

import (
	"time"

	"apps/api/app/domain"
	"github.com/google/uuid"
)

type ReleaseResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	ReleasedAt *time.Time `json:"released_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ReleaseSerializer struct{}

func (s ReleaseSerializer) ModelToResponse(release domain.ReleaseModel) ReleaseResponse {
	return ReleaseResponse{
		ID:         release.ID,
		Name:       release.Name,
		Status:     release.Status,
		ReleasedAt: release.ReleasedAt,
		CreatedAt:  release.CreatedAt,
		UpdatedAt:  release.UpdatedAt,
	}
}

func (s ReleaseSerializer) ModelToResponseMany(releases []domain.ReleaseModel) []ReleaseResponse {
	responses := make([]ReleaseResponse, len(releases))
	for i, release := range releases {
		responses[i] = s.ModelToResponse(release)
	}
	return responses
}
