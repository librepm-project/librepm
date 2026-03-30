package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockIssueAuditLogService struct {
	AllByIssueFunc           func(issueID uuid.UUID) (*[]domain.IssueAuditLogModel, error)
	LogFieldChangesFunc      func(issueID, userID uuid.UUID, old, new domain.IssueModel) error
	LogAttachmentAddedFunc   func(userID uuid.UUID, attachment domain.AttachmentModel) error
	LogAttachmentRemovedFunc func(userID uuid.UUID, attachment domain.AttachmentModel) error
	LogWorklogAddedFunc      func(userID uuid.UUID, worklog domain.IssueWorklogModel) error
	LogWorklogUpdatedFunc    func(userID uuid.UUID, old, new domain.IssueWorklogModel) error
	LogWorklogRemovedFunc    func(userID uuid.UUID, worklog domain.IssueWorklogModel) error
	LogRelatedIssueAddedFunc   func(issueID, userID uuid.UUID, relation domain.RelatedIssueModel, other domain.IssueModel) error
	LogRelatedIssueRemovedFunc func(issueID, userID uuid.UUID, relation domain.RelatedIssueModel, other domain.IssueModel) error
}

func (m *MockIssueAuditLogService) AllByIssue(issueID uuid.UUID) (*[]domain.IssueAuditLogModel, error) {
	if m.AllByIssueFunc != nil {
		return m.AllByIssueFunc(issueID)
	}
	return &[]domain.IssueAuditLogModel{}, nil
}

func (m *MockIssueAuditLogService) LogFieldChanges(issueID, userID uuid.UUID, old, new domain.IssueModel) error {
	if m.LogFieldChangesFunc != nil {
		return m.LogFieldChangesFunc(issueID, userID, old, new)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogAttachmentAdded(userID uuid.UUID, attachment domain.AttachmentModel) error {
	if m.LogAttachmentAddedFunc != nil {
		return m.LogAttachmentAddedFunc(userID, attachment)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogAttachmentRemoved(userID uuid.UUID, attachment domain.AttachmentModel) error {
	if m.LogAttachmentRemovedFunc != nil {
		return m.LogAttachmentRemovedFunc(userID, attachment)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogWorklogAdded(userID uuid.UUID, worklog domain.IssueWorklogModel) error {
	if m.LogWorklogAddedFunc != nil {
		return m.LogWorklogAddedFunc(userID, worklog)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogWorklogUpdated(userID uuid.UUID, old, new domain.IssueWorklogModel) error {
	if m.LogWorklogUpdatedFunc != nil {
		return m.LogWorklogUpdatedFunc(userID, old, new)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogWorklogRemoved(userID uuid.UUID, worklog domain.IssueWorklogModel) error {
	if m.LogWorklogRemovedFunc != nil {
		return m.LogWorklogRemovedFunc(userID, worklog)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogRelatedIssueAdded(issueID, userID uuid.UUID, relation domain.RelatedIssueModel, other domain.IssueModel) error {
	if m.LogRelatedIssueAddedFunc != nil {
		return m.LogRelatedIssueAddedFunc(issueID, userID, relation, other)
	}
	return nil
}

func (m *MockIssueAuditLogService) LogRelatedIssueRemoved(issueID, userID uuid.UUID, relation domain.RelatedIssueModel, other domain.IssueModel) error {
	if m.LogRelatedIssueRemovedFunc != nil {
		return m.LogRelatedIssueRemovedFunc(issueID, userID, relation, other)
	}
	return nil
}
