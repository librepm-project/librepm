package domain

import (
	"github.com/google/uuid"
)

type ProjectReleaseIssueServiceInterface interface {
	All() (*[]ProjectReleaseIssueModel, error)
	Show(id uuid.UUID) (*ProjectReleaseIssueModel, error)
	FindByProjectReleaseID(projectReleaseID uuid.UUID) (*[]ProjectReleaseIssueModel, error)
	FindByIssueID(issueID uuid.UUID) (*ProjectReleaseIssueModel, error)
	Create(pri *ProjectReleaseIssueModel) error
	Destroy(id uuid.UUID) error
}

type ProjectReleaseIssueService struct {
	ProjectReleaseIssueRepository ProjectReleaseIssueRepositoryInterface
}

func (s ProjectReleaseIssueService) All() (*[]ProjectReleaseIssueModel, error) {
	return s.ProjectReleaseIssueRepository.All()
}

func (s ProjectReleaseIssueService) Show(id uuid.UUID) (*ProjectReleaseIssueModel, error) {
	return s.ProjectReleaseIssueRepository.FindByID(id)
}

func (s ProjectReleaseIssueService) FindByProjectReleaseID(projectReleaseID uuid.UUID) (*[]ProjectReleaseIssueModel, error) {
	return s.ProjectReleaseIssueRepository.FindByProjectReleaseID(projectReleaseID)
}

func (s ProjectReleaseIssueService) FindByIssueID(issueID uuid.UUID) (*ProjectReleaseIssueModel, error) {
	return s.ProjectReleaseIssueRepository.FindByIssueID(issueID)
}

func (s ProjectReleaseIssueService) Create(pri *ProjectReleaseIssueModel) error {
	return s.ProjectReleaseIssueRepository.Create(pri)
}

func (s ProjectReleaseIssueService) Destroy(id uuid.UUID) error {
	return s.ProjectReleaseIssueRepository.Destroy(id)
}
