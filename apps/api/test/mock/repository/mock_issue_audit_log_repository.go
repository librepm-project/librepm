package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockIssueAuditLogRepository struct {
	AllByIssueFunc func(issueID uuid.UUID) (*[]domain.IssueAuditLogModel, error)
	CreateFunc     func(log *domain.IssueAuditLogModel) error
}

func (m *MockIssueAuditLogRepository) AllByIssue(issueID uuid.UUID) (*[]domain.IssueAuditLogModel, error) {
	if m.AllByIssueFunc != nil {
		return m.AllByIssueFunc(issueID)
	}
	return &[]domain.IssueAuditLogModel{}, nil
}

func (m *MockIssueAuditLogRepository) Create(log *domain.IssueAuditLogModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(log)
	}
	return nil
}
