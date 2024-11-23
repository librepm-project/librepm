package domain

import (
	"github.com/google/uuid"
)

type ProjectServiceInterface interface {
	All() *[]ProjectModel
	Show(project_id uuid.UUID) *ProjectModel
	Create(project *ProjectModel)
	Update(project_id uuid.UUID, project *ProjectModel)
	Destroy(project_id uuid.UUID)
}

type ProjectService struct {
	ProjectRepository ProjectRepositoryInterface
}

func (s ProjectService) All() *[]ProjectModel {
	projects := s.ProjectRepository.All()
	return projects
}

func (s ProjectService) Show(project_id uuid.UUID) *ProjectModel {
	return s.ProjectRepository.FindByID(project_id)
}

func (s ProjectService) Create(project *ProjectModel) {
	s.ProjectRepository.Create(project)
}

func (s ProjectService) Update(project_id uuid.UUID, project *ProjectModel) {
	s.ProjectRepository.Update(project_id, project)
}

func (s ProjectService) Destroy(project_id uuid.UUID) {
	s.ProjectRepository.Destroy(project_id)
}
