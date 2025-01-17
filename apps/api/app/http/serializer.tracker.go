package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type TrackerRequest struct {
	Name string `json:"name"`
}

type TrackerResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TrackerSerializer struct{}

func (s TrackerSerializer) RequestToModel(tracker_request TrackerRequest) domain.TrackerModel {
	return domain.TrackerModel{
		Name: tracker_request.Name,
	}
}

func (s TrackerSerializer) ModelToResponse(tracker domain.TrackerModel) TrackerResponse {
	return TrackerResponse{
		ID:   tracker.ID,
		Name: tracker.Name,
	}
}

func (s TrackerSerializer) ModelToResponseMany(trackers []domain.TrackerModel) []TrackerResponse {
	serialized := []TrackerResponse{}
	for _, tracker := range trackers {
		serialized = append(serialized, s.ModelToResponse(tracker))
	}
	return serialized
}
