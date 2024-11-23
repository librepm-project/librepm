package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type Project struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CodeName string    `json:"codeName"`
}

type ProjectSerializer struct{}

func (s ProjectSerializer) SerializeProject(project domain.ProjectModel) Project {
	return Project{
		ID:       project.ID,
		Name:     project.Name,
		CodeName: project.CodeName,
	}
}

func (s ProjectSerializer) SerializeProjects(projects []domain.ProjectModel) []Project {
	var serialized []Project
	for _, project := range projects {
		serialized = append(serialized, s.SerializeProject(project))
	}
	return serialized
}
