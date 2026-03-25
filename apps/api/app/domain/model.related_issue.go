package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RelatedIssueModel struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key;"`
	IssueAID  uuid.UUID  `gorm:"type:char(36);not null;index"`
	IssueBID  uuid.UUID  `gorm:"type:char(36);not null;index"`
	Type      string     `gorm:"type:varchar(20);not null;"` // 'related', 'blocks', 'implements', 'duplicates'
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	IssueA IssueModel `gorm:"foreignKey:IssueAID"`
	IssueB IssueModel `gorm:"foreignKey:IssueBID"`
}

func (r RelatedIssueModel) TableName() string {
	return "related_issue"
}

func (r *RelatedIssueModel) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
