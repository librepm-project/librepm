package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	AuditEventFieldChanged        = "field_changed"
	AuditEventAttachmentAdded     = "attachment_added"
	AuditEventAttachmentRemoved   = "attachment_removed"
	AuditEventWorklogAdded        = "worklog_added"
	AuditEventWorklogUpdated      = "worklog_updated"
	AuditEventWorklogRemoved      = "worklog_removed"
	AuditEventRelatedIssueAdded   = "related_issue_added"
	AuditEventRelatedIssueRemoved = "related_issue_removed"
)

type IssueAuditLogModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	IssueID   uuid.UUID `gorm:"type:char(36);not null;index;"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;"`
	EventType string    `gorm:"type:varchar(50);not null;"`
	Field     string    `gorm:"type:varchar(100);"`
	OldValue  string    `gorm:"type:text;"`
	NewValue  string    `gorm:"type:text;"`
	Meta      string    `gorm:"type:text;"`
	CreatedAt time.Time

	User UserModel `gorm:"foreignKey:UserID"`
}

func (l *IssueAuditLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
