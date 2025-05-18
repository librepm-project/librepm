package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type IssueServiceInterface interface {
	All() (*[]IssueModel, error)
	Show(issue_id uuid.UUID) (*IssueModel, error)
	Create(issue *IssueModel) error
	Update(issue_id uuid.UUID, issue *IssueModel) error
	Destroy(issue_id uuid.UUID) error
}

type IssueService struct {
	IssueRepository   IssueRepositoryInterface
	ProjectRepository ProjectRepositoryInterface
}

func (s IssueService) All() (*[]IssueModel, error) {
	issues, err := s.IssueRepository.All()
	return issues, err
}

func (s IssueService) Show(issue_id uuid.UUID) (*IssueModel, error) {
	return s.IssueRepository.FindByID(issue_id)
}

func (s IssueService) Create(issue *IssueModel) error {
	err := s.ProjectRepository.IncrementLastIssueKeyNumber(issue.ProjectID)
	if err != nil {
		return err
	}
	project, err := s.ProjectRepository.FindByID(issue.ProjectID)
	if err != nil {
		return err
	}
	issue.Key = project.CodeName + "-" + fmt.Sprint(project.LastIssueKeyNumber)
	return s.IssueRepository.Create(issue)
}

func (s IssueService) Update(issue_id uuid.UUID, issue *IssueModel) error {
	return s.IssueRepository.Update(issue_id, issue)
}

func (s IssueService) Destroy(issue_id uuid.UUID) error {
	return s.IssueRepository.Destroy(issue_id)
}
