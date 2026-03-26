package domain

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type IssueAuditLogServiceInterface interface {
	AllByIssue(issueID uuid.UUID) (*[]IssueAuditLogModel, error)
	LogFieldChanges(issueID, userID uuid.UUID, old, new IssueModel) error
	LogAttachmentAdded(userID uuid.UUID, attachment AttachmentModel) error
	LogAttachmentRemoved(userID uuid.UUID, attachment AttachmentModel) error
	LogWorklogAdded(userID uuid.UUID, worklog IssueWorklogModel) error
	LogWorklogUpdated(userID uuid.UUID, old, new IssueWorklogModel) error
	LogWorklogRemoved(userID uuid.UUID, worklog IssueWorklogModel) error
	LogRelatedIssueAdded(issueID, userID uuid.UUID, relation RelatedIssueModel, other IssueModel) error
	LogRelatedIssueRemoved(issueID, userID uuid.UUID, relation RelatedIssueModel, other IssueModel) error
}

type IssueAuditLogService struct {
	IssueAuditLogRepository IssueAuditLogRepositoryInterface
}

func (s IssueAuditLogService) AllByIssue(issueID uuid.UUID) (*[]IssueAuditLogModel, error) {
	return s.IssueAuditLogRepository.AllByIssue(issueID)
}

// logEntry creates and persists an audit log entry
func (s IssueAuditLogService) logEntry(issueID, userID uuid.UUID, eventType, oldValue, newValue string, meta map[string]interface{}) error {
	var metaStr string
	if meta != nil {
		metaBytes, _ := json.Marshal(meta)
		metaStr = string(metaBytes)
	}
	entry := IssueAuditLogModel{
		IssueID:   issueID,
		UserID:    userID,
		EventType: eventType,
		OldValue:  oldValue,
		NewValue:  newValue,
		Meta:      metaStr,
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogFieldChanges(issueID, userID uuid.UUID, old, new IssueModel) error {
	type fieldDiff struct {
		name     string
		oldValue string
		newValue string
	}

	diffs := []fieldDiff{}

	if old.Summary != new.Summary {
		diffs = append(diffs, fieldDiff{"summary", old.Summary, new.Summary})
	}
	if old.Description != new.Description {
		diffs = append(diffs, fieldDiff{"description", old.Description, new.Description})
	}
	if old.StatusID != new.StatusID {
		diffs = append(diffs, fieldDiff{"status", old.Status.Name, new.Status.Name})
	}
	if old.TrackerID != new.TrackerID {
		diffs = append(diffs, fieldDiff{"tracker", old.Tracker.Name, new.Tracker.Name})
	}
	oldParent := ""
	newParent := ""
	if old.ParentID != nil {
		oldParent = old.ParentID.String()
	}
	if new.ParentID != nil {
		newParent = new.ParentID.String()
	}
	if oldParent != newParent {
		diffs = append(diffs, fieldDiff{"parent", oldParent, newParent})
	}

	for _, diff := range diffs {
		if err := s.logEntry(issueID, userID, AuditEventFieldChanged, diff.oldValue, diff.newValue, map[string]interface{}{"field": diff.name}); err != nil {
			return err
		}
	}
	return nil
}

func (s IssueAuditLogService) LogAttachmentAdded(userID uuid.UUID, attachment AttachmentModel) error {
	return s.logEntry(
		attachment.EntityID,
		userID,
		AuditEventAttachmentAdded,
		"",
		attachment.Filename,
		map[string]interface{}{
			"filename": attachment.Filename,
			"size":     attachment.Size,
		},
	)
}

func (s IssueAuditLogService) LogAttachmentRemoved(userID uuid.UUID, attachment AttachmentModel) error {
	return s.logEntry(
		attachment.EntityID,
		userID,
		AuditEventAttachmentRemoved,
		attachment.Filename,
		"",
		map[string]interface{}{
			"filename": attachment.Filename,
			"size":     attachment.Size,
		},
	)
}

func (s IssueAuditLogService) LogWorklogAdded(userID uuid.UUID, worklog IssueWorklogModel) error {
	return s.logEntry(
		worklog.IssueID,
		userID,
		AuditEventWorklogAdded,
		"",
		fmt.Sprint(worklog.Seconds),
		map[string]interface{}{
			"comment": worklog.Comment,
		},
	)
}

func (s IssueAuditLogService) LogWorklogUpdated(userID uuid.UUID, old, new IssueWorklogModel) error {
	return s.logEntry(
		new.IssueID,
		userID,
		AuditEventWorklogUpdated,
		fmt.Sprint(old.Seconds),
		fmt.Sprint(new.Seconds),
		map[string]interface{}{
			"comment": new.Comment,
		},
	)
}

func (s IssueAuditLogService) LogWorklogRemoved(userID uuid.UUID, worklog IssueWorklogModel) error {
	return s.logEntry(
		worklog.IssueID,
		userID,
		AuditEventWorklogRemoved,
		fmt.Sprint(worklog.Seconds),
		"",
		map[string]interface{}{
			"comment": worklog.Comment,
		},
	)
}

func (s IssueAuditLogService) LogRelatedIssueAdded(issueID, userID uuid.UUID, relation RelatedIssueModel, other IssueModel) error {
	return s.logEntry(
		issueID,
		userID,
		AuditEventRelatedIssueAdded,
		"",
		relation.Type,
		map[string]interface{}{
			"issue_key":     other.Key,
			"issue_summary": other.Summary,
			"relation_type": relation.Type,
		},
	)
}

func (s IssueAuditLogService) LogRelatedIssueRemoved(issueID, userID uuid.UUID, relation RelatedIssueModel, other IssueModel) error {
	return s.logEntry(
		issueID,
		userID,
		AuditEventRelatedIssueRemoved,
		relation.Type,
		"",
		map[string]interface{}{
			"issue_key":     other.Key,
			"issue_summary": other.Summary,
			"relation_type": relation.Type,
		},
	)
}
