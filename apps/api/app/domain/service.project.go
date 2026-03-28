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
	ShowIssuePropertyById(project_id uuid.UUID) (*ProjectIssueProperty, error)
	TrackersByProject(project_id uuid.UUID) (*[]TrackerModel, error)
	StatusesByProject(project_id uuid.UUID) (*[]StatusModel, error)
}

type ProjectService struct {
	ProjectRepository              ProjectRepositoryInterface
	ProjectIssuePropertyRepository ProjectIssuePropertyRepositoryInterface
	ProjectTrackerRepository       ProjectTrackerRepositoryInterface
	ProjectTrackerStatusRepository ProjectTrackerStatusRepositoryInterface
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

func (s ProjectService) ShowIssuePropertyById(project_id uuid.UUID) (*ProjectIssueProperty, error) {
	return s.ProjectIssuePropertyRepository.FindByProjectId(project_id)
}

func (s ProjectService) TrackersByProject(project_id uuid.UUID) (*[]TrackerModel, error) {
	return s.ProjectTrackerRepository.FindTrackersByProjectID(project_id)
}

func (s ProjectService) StatusesByProject(project_id uuid.UUID) (*[]StatusModel, error) {
	return s.ProjectTrackerStatusRepository.FindStatusesByProjectID(project_id)
}
