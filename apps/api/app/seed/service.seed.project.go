package seed

import (
	"apps/api/app/domain"
	"fmt"
	"time"
)

func (s SeedService) createProject(items []ProjectData) error {
	var err error
	var project domain.ProjectModel

	for _, item := range items {
		project = domain.ProjectModel{
			Name:      item.Name,
			CodeName:  item.CodeName,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = s.ProjectRepository.Create(&project)
		fmt.Println(err)

		for _, tracker_data := range item.Trackers {

			tracker, _ := s.TrackerRepository.FindByName(tracker_data.Name)

			project_tracker := domain.ProjectTrackerModel{
				ProjectID: project.ID,
				TrackerID: tracker.ID,
			}
			err = s.ProjectTrackerRepository.Create(&project_tracker)

			for _, statusName := range tracker_data.Statuses {

				status, _ := s.StatusRepository.FindByName(statusName)

				err = s.ProjectTrackerStatusRepository.Create(&domain.ProjectTrackerStatusModel{
					StatusID:         status.ID,
					ProjectTrackerID: project_tracker.ID,
				})
			}
		}
	}
	return err
}
