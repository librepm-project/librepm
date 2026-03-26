package domain

import (
	"github.com/google/uuid"
)

type IssueWorklogServiceInterface interface {
	AllByIssue(issueId uuid.UUID) (*[]IssueWorklogModel, error)
	Show(worklogId uuid.UUID) (*IssueWorklogModel, error)
	Create(worklog *IssueWorklogModel) error
	Update(worklogId uuid.UUID, worklog *IssueWorklogModel) error
	Destroy(worklogId uuid.UUID) error
}

type IssueWorklogService struct {
	IssueWorklogRepository IssueWorklogRepositoryInterface
}

func (s IssueWorklogService) AllByIssue(issueId uuid.UUID) (*[]IssueWorklogModel, error) {
	return s.IssueWorklogRepository.AllByIssue(issueId)
}

func (s IssueWorklogService) Show(worklogId uuid.UUID) (*IssueWorklogModel, error) {
	return s.IssueWorklogRepository.FindByID(worklogId)
}

func (s IssueWorklogService) Create(worklog *IssueWorklogModel) error {
	return s.IssueWorklogRepository.Create(worklog)
}

func (s IssueWorklogService) Update(worklogId uuid.UUID, worklog *IssueWorklogModel) error {
	return s.IssueWorklogRepository.Update(worklogId, worklog)
}

func (s IssueWorklogService) Destroy(worklogId uuid.UUID) error {
	return s.IssueWorklogRepository.Destroy(worklogId)
}
