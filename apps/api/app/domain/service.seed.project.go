package domain

import (
	"time"
)

type ProjectData struct {
	Name     string `yaml:"name"`
	CodeName string `yaml:"code_name"`
}

func (s SeedService) createProject(items []ProjectData) error {
	var err error
	var project ProjectModel
	for _, item := range items {
		project = ProjectModel{
			Name:      item.Name,
			CodeName:  item.CodeName,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = s.ProjectRepository.Create(&project)
	}
	return err
}
