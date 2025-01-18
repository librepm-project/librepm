package domain

import (
	"github.com/google/uuid"
)

type ProjectServiceInterface interface {
	All() (*[]ProjectModel, error)
	Show(project_id uuid.UUID) (*ProjectModel, error)
	Create(project *ProjectModel) error
	Update(project_id uuid.UUID, project *ProjectModel) error
	Destroy(project_id uuid.UUID) error
}

type ProjectService struct {
	ProjectRepository ProjectRepositoryInterface
}

func (s ProjectService) All() (*[]ProjectModel, error) {
	projects, err := s.ProjectRepository.All()
	return projects, err
}

func (s ProjectService) Show(project_id uuid.UUID) (*ProjectModel, error) {
	return s.ProjectRepository.FindByID(project_id)
}

func (s ProjectService) Create(project *ProjectModel) error {
	return s.ProjectRepository.Create(project)
}

func (s ProjectService) Update(project_id uuid.UUID, project *ProjectModel) error {
	return s.ProjectRepository.Update(project_id, project)
}

func (s ProjectService) Destroy(project_id uuid.UUID) error {
	return s.ProjectRepository.Destroy(project_id)
}
