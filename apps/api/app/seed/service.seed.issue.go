package seed

import (
	"apps/api/app/domain"
	"fmt"
)

func (s SeedService) createIssue(items []IssueData) error {
	var err error
	records := make(map[string]*domain.IssueModel)

	for i, item := range items {
		status, _ := s.StatusRepository.FindByName(item.StatusName)
		project, _ := s.ProjectRepository.FindByName(item.ProjectName)
		tracker, _ := s.TrackerRepository.FindByName(item.TrackerName)

		issue := &domain.IssueModel{
			Summary:     item.Summary,
			Key:         fmt.Sprintf("%s-%d", project.CodeName, i),
			Description: item.Description,
			StatusID:    status.ID,
			ProjectID:   project.ID,
			TrackerID:   tracker.ID,
		}

		if item.ParentName != "" {
			if parent, ok := records[item.ParentName]; ok {
				issue.ParentID = &parent.ID
			}
		}

		if err = s.IssueRepository.Create(issue); err != nil {
			return fmt.Errorf("failed to create issue '%s': %w", item.Summary, err)
		}

		records[item.Summary] = issue
	}

	return nil
}
