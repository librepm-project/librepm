package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueModel struct {
	ID          uuid.UUID  `gorm:"type:char(36);primary_key;"`
	ProjectID   uuid.UUID  `gorm:"type:char(36);not null;uniqueIndex:project_id_key_summary"`
	TrackerID   uuid.UUID  `gorm:"type:char(36) references tracker;not null;"`
	StatusID    uuid.UUID  `gorm:"type:char(36) references status;not null;"`
	ParentID    *uuid.UUID `gorm:"type:char(36) references issue;uniqueIndex:project_id_key_summary"`
	Key         string     `gorm:"type:varchar(10);not null;"`
	Summary     string     `gorm:"type:varchar(100);not null;uniqueIndex:project_id_key_summary"`
	Description string     `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	AssignedUserID  *uuid.UUID `gorm:"type:char(36) references user;"`
	ReporterUserID  *uuid.UUID `gorm:"type:char(36) references user;"`
	PriorityID      *uuid.UUID `gorm:"type:char(36) references priority;"`

	Project      ProjectModel
	Tracker      TrackerModel
	Status       StatusModel
	AssignedUser *UserModel
	ReporterUser *UserModel
	Priority     *PriorityModel

	Parent *IssueModel
}

func (issue IssueModel) TableName() string {
	return "issue"
}

func (issue *IssueModel) BeforeCreate(tx *gorm.DB) (err error) {
	issue.ID = uuid.New()
	return
}
