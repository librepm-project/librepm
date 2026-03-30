package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockIssueWorklogService struct {
	AllByIssueFunc func(issueID uuid.UUID) (*[]domain.IssueWorklogModel, error)
	ShowFunc       func(worklogID uuid.UUID) (*domain.IssueWorklogModel, error)
	CreateFunc     func(worklog *domain.IssueWorklogModel) error
	UpdateFunc     func(worklogID uuid.UUID, worklog *domain.IssueWorklogModel) error
	DestroyFunc    func(worklogID uuid.UUID) error
}

func (m *MockIssueWorklogService) AllByIssue(issueID uuid.UUID) (*[]domain.IssueWorklogModel, error) {
	if m.AllByIssueFunc != nil {
		return m.AllByIssueFunc(issueID)
	}
	return &[]domain.IssueWorklogModel{}, nil
}

func (m *MockIssueWorklogService) Show(worklogID uuid.UUID) (*domain.IssueWorklogModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(worklogID)
	}
	return &domain.IssueWorklogModel{}, nil
}

func (m *MockIssueWorklogService) Create(worklog *domain.IssueWorklogModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(worklog)
	}
	return nil
}

func (m *MockIssueWorklogService) Update(worklogID uuid.UUID, worklog *domain.IssueWorklogModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(worklogID, worklog)
	}
	return nil
}

func (m *MockIssueWorklogService) Destroy(worklogID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(worklogID)
	}
	return nil
}
