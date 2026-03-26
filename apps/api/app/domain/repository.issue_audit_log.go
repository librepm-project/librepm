package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueAuditLogRepositoryInterface interface {
	AllByIssue(issueID uuid.UUID) (*[]IssueAuditLogModel, error)
	Create(log *IssueAuditLogModel) error
}

type IssueAuditLogRepository struct {
	DB *gorm.DB
}

func (r IssueAuditLogRepository) AllByIssue(issueID uuid.UUID) (*[]IssueAuditLogModel, error) {
	var logs []IssueAuditLogModel
	err := r.DB.Preload("User").Where("issue_id = ?", issueID).Order("created_at DESC").Find(&logs).Error
	if err != nil {
		slog.Error("failed to fetch audit logs", "error", err)
	}
	return &logs, err
}

func (r IssueAuditLogRepository) Create(log *IssueAuditLogModel) error {
	err := r.DB.Create(log).Error
	if err != nil {
		slog.Error("failed to create audit log", "error", err)
	}
	return err
}
