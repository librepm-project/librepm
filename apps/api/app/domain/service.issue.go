package domain

import (
	"github.com/google/uuid"
)

type IssueServiceInterface interface {
	All() *[]IssueModel
	Show(issue_id uuid.UUID) *IssueModel
	Create(issue *IssueModel)
	Update(issue_id uuid.UUID, issue *IssueModel)
	Destroy(issue_id uuid.UUID)
}

type IssueService struct {
	IssueRepository IssueRepositoryInterface
}

func (s IssueService) All() *[]IssueModel {
	issues := s.IssueRepository.All()
	return issues
}

func (s IssueService) Show(issue_id uuid.UUID) *IssueModel {
	return s.IssueRepository.FindByID(issue_id)
}

func (s IssueService) Create(issue *IssueModel) {
	s.IssueRepository.Create(issue)
}

func (s IssueService) Update(issue_id uuid.UUID, issue *IssueModel) {
	s.IssueRepository.Update(issue_id, issue)
}

func (s IssueService) Destroy(issue_id uuid.UUID) {
	s.IssueRepository.Destroy(issue_id)
}
