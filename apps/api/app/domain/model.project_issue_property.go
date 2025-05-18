package domain

import "github.com/google/uuid"

type ProjectIssuePropertyTrackerStatus struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ProjectIssuePropertyTracker struct {
	ID       uuid.UUID                           `json:"id"`
	Name     string                              `json:"name"`
	Statuses []ProjectIssuePropertyTrackerStatus `json:"statuses"`
}

type ProjectIssueProperty struct {
	Trackers []ProjectIssuePropertyTracker `json:"trackers"`
}

type ProjectIssuePropertyQueryResult struct {
	TrackerID   uuid.UUID
	TrackerName string
	StatusID    uuid.UUID
	StatusName  string
}
