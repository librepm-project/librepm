package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectIssuePropertyRepositoryInterface interface {
	FindByProjectId(project_id uuid.UUID) (*ProjectIssueProperty, error)
}

type ProjectIssuePropertyRepository struct {
	DB *gorm.DB
}

func (r ProjectIssuePropertyRepository) FindByProjectId(project_id uuid.UUID) (*ProjectIssueProperty, error) {

	var result []ProjectIssuePropertyQueryResult

	err := r.DB.
		Table("project").
		Select("tracker.id as tracker_id, tracker.name as tracker_name, status.id as status_id, status.name as status_name").
		Joins("JOIN project_tracker ON project.id = project_tracker.project_id").
		Joins("JOIN tracker ON project_tracker.tracker_id = tracker.id").
		Joins("JOIN project_tracker_status ON project_tracker_status.project_tracker_id = project_tracker.id").
		Joins("JOIN status ON project_tracker_status.status_id = status.id").
		Where("project.id = ?", project_id).
		Scan(&result).Error

	if err != nil {
		fmt.Println(err.Error())
	}

	project_issue_property := &ProjectIssueProperty{}

	trackers := make(map[uuid.UUID]*ProjectIssuePropertyTracker)
	tracker_ids := []uuid.UUID{}

	for _, row := range result {
		tracker, exists := trackers[row.TrackerID]
		if !exists {
			tracker = &ProjectIssuePropertyTracker{
				ID:       row.TrackerID,
				Name:     row.TrackerName,
				Statuses: []ProjectIssuePropertyTrackerStatus{},
			}
			trackers[row.TrackerID] = tracker
			tracker_ids = append(tracker_ids, row.TrackerID)
		}

		tracker.Statuses = append(tracker.Statuses, ProjectIssuePropertyTrackerStatus{
			ID:   row.StatusID,
			Name: row.StatusName,
		})
	}

	project_issue_property.Trackers = make([]ProjectIssuePropertyTracker, 0, len(trackers))
	for _, id := range tracker_ids {
		if tracker, ok := trackers[id]; ok {
			project_issue_property.Trackers = append(project_issue_property.Trackers, *tracker)
		}
	}

	return project_issue_property, err
}
