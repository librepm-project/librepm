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
		entry := IssueAuditLogModel{
			IssueID:   issueID,
			UserID:    userID,
			EventType: AuditEventFieldChanged,
			Field:     diff.name,
			OldValue:  diff.oldValue,
			NewValue:  diff.newValue,
		}
		if err := s.IssueAuditLogRepository.Create(&entry); err != nil {
			return err
		}
	}
	return nil
}

func (s IssueAuditLogService) LogAttachmentAdded(userID uuid.UUID, attachment AttachmentModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"filename": attachment.Filename,
		"size":     attachment.Size,
	})
	entry := IssueAuditLogModel{
		IssueID:   attachment.EntityID,
		UserID:    userID,
		EventType: AuditEventAttachmentAdded,
		NewValue:  attachment.Filename,
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogAttachmentRemoved(userID uuid.UUID, attachment AttachmentModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"filename": attachment.Filename,
		"size":     attachment.Size,
	})
	entry := IssueAuditLogModel{
		IssueID:   attachment.EntityID,
		UserID:    userID,
		EventType: AuditEventAttachmentRemoved,
		OldValue:  attachment.Filename,
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogWorklogAdded(userID uuid.UUID, worklog IssueWorklogModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"comment": worklog.Comment,
	})
	entry := IssueAuditLogModel{
		IssueID:   worklog.IssueID,
		UserID:    userID,
		EventType: AuditEventWorklogAdded,
		NewValue:  fmt.Sprint(worklog.Seconds),
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogWorklogUpdated(userID uuid.UUID, old, new IssueWorklogModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"comment": new.Comment,
	})
	entry := IssueAuditLogModel{
		IssueID:   new.IssueID,
		UserID:    userID,
		EventType: AuditEventWorklogUpdated,
		OldValue:  fmt.Sprint(old.Seconds),
		NewValue:  fmt.Sprint(new.Seconds),
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogWorklogRemoved(userID uuid.UUID, worklog IssueWorklogModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"comment": worklog.Comment,
	})
	entry := IssueAuditLogModel{
		IssueID:   worklog.IssueID,
		UserID:    userID,
		EventType: AuditEventWorklogRemoved,
		OldValue:  fmt.Sprint(worklog.Seconds),
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogRelatedIssueAdded(issueID, userID uuid.UUID, relation RelatedIssueModel, other IssueModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"issue_key":     other.Key,
		"issue_summary": other.Summary,
		"relation_type": relation.Type,
	})
	entry := IssueAuditLogModel{
		IssueID:   issueID,
		UserID:    userID,
		EventType: AuditEventRelatedIssueAdded,
		NewValue:  relation.Type,
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}

func (s IssueAuditLogService) LogRelatedIssueRemoved(issueID, userID uuid.UUID, relation RelatedIssueModel, other IssueModel) error {
	meta, _ := json.Marshal(map[string]interface{}{
		"issue_key":     other.Key,
		"issue_summary": other.Summary,
		"relation_type": relation.Type,
	})
	entry := IssueAuditLogModel{
		IssueID:   issueID,
		UserID:    userID,
		EventType: AuditEventRelatedIssueRemoved,
		OldValue:  relation.Type,
		Meta:      string(meta),
	}
	return s.IssueAuditLogRepository.Create(&entry)
}
