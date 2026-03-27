package domain

import (
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueRepositoryInterface interface {
	All() (*[]IssueModel, error)
	AllByFilter(conditions []FilterConditionModel) (*[]IssueModel, error)
	FindByID(issue_id uuid.UUID) (*IssueModel, error)
	FindByKey(key string) (*IssueModel, error)
	Create(issue *IssueModel) error
	Update(issue_id uuid.UUID, issue *IssueModel) error
	Destroy(issue_id uuid.UUID) error
}

type IssueRepository struct {
	DB *gorm.DB
}

func (r IssueRepository) withAssociations(query *gorm.DB) *gorm.DB {
	return query.Preload("Project").Preload("Tracker").Preload("Status").Preload("AssignedUser").Preload("ReporterUser").Preload("Priority")
}

func (r IssueRepository) All() (*[]IssueModel, error) {
	var issues []IssueModel
	var err error
	query := r.DB.Select("issue.*")

	if err := r.withAssociations(query).Find(&issues).Error; err != nil {
		slog.Error("failed to fetch issues", "error", err)
	}
	return &issues, err
}

func (r IssueRepository) AllByFilter(conditions []FilterConditionModel) (*[]IssueModel, error) {
	var issues []IssueModel
	query := r.DB.Select("issue.*")

	for _, c := range conditions {
		switch c.Field {
		case "project_id":
			if c.Op == "eq" {
				query = query.Where("issue.project_id = ?", c.Value)
			} else if c.Op == "ne" {
				query = query.Where("issue.project_id != ?", c.Value)
			}
		case "tracker_id":
			if c.Op == "eq" {
				query = query.Where("issue.tracker_id = ?", c.Value)
			} else if c.Op == "ne" {
				query = query.Where("issue.tracker_id != ?", c.Value)
			}
		case "status_id":
			if c.Op == "eq" {
				query = query.Where("issue.status_id = ?", c.Value)
			} else if c.Op == "ne" {
				query = query.Where("issue.status_id != ?", c.Value)
			}
		case "summary":
			if c.Op == "contains" {
				query = query.Where("issue.summary LIKE ?", "%"+strings.TrimSpace(c.Value)+"%")
			} else if c.Op == "eq" {
				query = query.Where("issue.summary = ?", c.Value)
			}
		case "assigned_user_id":
			if c.Op == "eq" {
				query = query.Where("issue.assigned_user_id = ?", c.Value)
			} else if c.Op == "ne" {
				query = query.Where("issue.assigned_user_id != ?", c.Value)
			}
		case "reporter_user_id":
			if c.Op == "eq" {
				query = query.Where("issue.reporter_user_id = ?", c.Value)
			} else if c.Op == "ne" {
				query = query.Where("issue.reporter_user_id != ?", c.Value)
			}
		case "priority_id":
			if c.Op == "eq" {
				query = query.Where("issue.priority_id = ?", c.Value)
			} else if c.Op == "ne" {
				query = query.Where("issue.priority_id != ?", c.Value)
			}
		}
	}

	if err := r.withAssociations(query).Find(&issues).Error; err != nil {
		slog.Error("failed to fetch issues by filter", "error", err)
		return nil, err
	}
	return &issues, nil
}

func (r IssueRepository) FindByID(issue_id uuid.UUID) (*IssueModel, error) {
	var issue IssueModel
	query := r.withAssociations(r.DB).First(&issue, issue_id)

	if query.Error != nil {
		slog.Error("failed to find issue by id", "error", query.Error)
	}
	return &issue, query.Error
}

func (r IssueRepository) FindByKey(key string) (*IssueModel, error) {
	var issue IssueModel
	query := r.withAssociations(r.DB).Where("`key` = ?", key).First(&issue)

	if query.Error != nil {
		slog.Error("failed to find issue by key", "error", query.Error)
	}
	return &issue, query.Error
}

func (r IssueRepository) Create(issue *IssueModel) error {
	query := r.DB.Create(&issue)
	if query.Error != nil {
		slog.Error("failed to create issue", "error", query.Error)
	}
	return query.Error
}

func (r IssueRepository) Update(issue_id uuid.UUID, issue *IssueModel) error {
	query := r.DB.Model(IssueModel{}).Where("id", issue_id).Updates(&issue)
	if query.Error != nil {
		slog.Error("failed to update issue", "error", query.Error)
	}
	return query.Error
}

func (r IssueRepository) Destroy(issue_id uuid.UUID) error {
	query := r.DB.Delete(&IssueModel{}, issue_id)
	if query.Error != nil {
		slog.Error("failed to destroy issue", "error", query.Error)
	}
	return query.Error
}
