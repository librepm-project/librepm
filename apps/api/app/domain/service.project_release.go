package domain

import (
	"github.com/google/uuid"
)

type ProjectReleaseServiceInterface interface {
	All() (*[]ProjectReleaseModel, error)
	Show(id uuid.UUID) (*ProjectReleaseModel, error)
	FindByReleaseID(releaseID uuid.UUID) (*[]ProjectReleaseModel, error)
	FindByProjectID(projectID uuid.UUID) (*[]ProjectReleaseModel, error)
	Create(projectRelease *ProjectReleaseModel) error
	Update(id uuid.UUID, projectRelease *ProjectReleaseModel) error
	Destroy(id uuid.UUID) error
}

type ProjectReleaseService struct {
	ProjectReleaseRepository ProjectReleaseRepositoryInterface
}

func (s ProjectReleaseService) All() (*[]ProjectReleaseModel, error) {
	return s.ProjectReleaseRepository.All()
}

func (s ProjectReleaseService) Show(id uuid.UUID) (*ProjectReleaseModel, error) {
	return s.ProjectReleaseRepository.FindByID(id)
}

func (s ProjectReleaseService) FindByReleaseID(releaseID uuid.UUID) (*[]ProjectReleaseModel, error) {
	return s.ProjectReleaseRepository.FindByReleaseID(releaseID)
}

func (s ProjectReleaseService) FindByProjectID(projectID uuid.UUID) (*[]ProjectReleaseModel, error) {
	return s.ProjectReleaseRepository.FindByProjectID(projectID)
}

func (s ProjectReleaseService) Create(projectRelease *ProjectReleaseModel) error {
	return s.ProjectReleaseRepository.Create(projectRelease)
}

func (s ProjectReleaseService) Update(id uuid.UUID, projectRelease *ProjectReleaseModel) error {
	return s.ProjectReleaseRepository.Update(id, projectRelease)
}

func (s ProjectReleaseService) Destroy(id uuid.UUID) error {
	return s.ProjectReleaseRepository.Destroy(id)
}
