package seed

import (
	"apps/api/app/domain"
	"time"
)

type ProjectData struct {
	Name     string   `yaml:"name"`
	CodeName string   `yaml:"code_name"`
	Trackers []string `yaml:"trackers"`
}

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

		for _, trackerName := range item.Trackers {

			tracker, _ := s.TrackerRepository.FindByName(trackerName)

			err = s.ProjectTrackerRepository.Create(&domain.ProjectTrackerModel{
				ProjectID: project.ID,
				TrackerID: tracker.ID,
			})
		}
	}
	return err
}
