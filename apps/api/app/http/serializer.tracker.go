package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Tracker struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TrackerSerializer struct{}

func (s TrackerSerializer) SerializeTracker(tracker domain.TrackerModel) Tracker {
	return Tracker{
		ID:   tracker.ID,
		Name: tracker.Name,
	}
}

func (s TrackerSerializer) SerializeTrackers(trackers []domain.TrackerModel) []Tracker {
	var serialized []Tracker
	for _, tracker := range trackers {
		serialized = append(serialized, s.SerializeTracker(tracker))
	}
	return serialized
}
