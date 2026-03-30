package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockIssueWorklogRepository struct {
	AllByIssueFunc func(issueId uuid.UUID) (*[]domain.IssueWorklogModel, error)
	FindByIDFunc   func(worklogId uuid.UUID) (*domain.IssueWorklogModel, error)
	CreateFunc     func(worklog *domain.IssueWorklogModel) error
	UpdateFunc     func(worklogId uuid.UUID, worklog *domain.IssueWorklogModel) error
	DestroyFunc    func(worklogId uuid.UUID) error
}

func (m *MockIssueWorklogRepository) AllByIssue(issueId uuid.UUID) (*[]domain.IssueWorklogModel, error) {
	if m.AllByIssueFunc != nil {
		return m.AllByIssueFunc(issueId)
	}
	return &[]domain.IssueWorklogModel{}, nil
}

func (m *MockIssueWorklogRepository) FindByID(worklogId uuid.UUID) (*domain.IssueWorklogModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(worklogId)
	}
	return &domain.IssueWorklogModel{}, nil
}

func (m *MockIssueWorklogRepository) Create(worklog *domain.IssueWorklogModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(worklog)
	}
	return nil
}

func (m *MockIssueWorklogRepository) Update(worklogId uuid.UUID, worklog *domain.IssueWorklogModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(worklogId, worklog)
	}
	return nil
}

func (m *MockIssueWorklogRepository) Destroy(worklogId uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(worklogId)
	}
	return nil
}
