package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueWorklogRepositoryInterface interface {
	AllByIssue(issueId uuid.UUID) (*[]IssueWorklogModel, error)
	FindByID(worklogId uuid.UUID) (*IssueWorklogModel, error)
	Create(worklog *IssueWorklogModel) error
	Update(worklogId uuid.UUID, worklog *IssueWorklogModel) error
	Destroy(worklogId uuid.UUID) error
}

type IssueWorklogRepository struct {
	DB *gorm.DB
}

func (r IssueWorklogRepository) AllByIssue(issueId uuid.UUID) (*[]IssueWorklogModel, error) {
	var worklogs []IssueWorklogModel
	err := r.DB.Preload("User").Where("issue_id = ?", issueId).Order("logged_at DESC").Find(&worklogs).Error
	if err != nil {
		slog.Error("failed to fetch worklogs", "error", err)
	}
	return &worklogs, err
}

func (r IssueWorklogRepository) FindByID(worklogId uuid.UUID) (*IssueWorklogModel, error) {
	var worklog IssueWorklogModel
	err := r.DB.Preload("User").First(&worklog, worklogId).Error
	if err != nil {
		slog.Error("failed to find worklog by id", "error", err)
	}
	return &worklog, err
}

func (r IssueWorklogRepository) Create(worklog *IssueWorklogModel) error {
	err := r.DB.Create(worklog).Error
	if err != nil {
		slog.Error("failed to create worklog", "error", err)
	}
	return err
}

func (r IssueWorklogRepository) Update(worklogId uuid.UUID, worklog *IssueWorklogModel) error {
	err := r.DB.Model(IssueWorklogModel{}).Where("id = ?", worklogId).Updates(worklog).Error
	if err != nil {
		slog.Error("failed to update worklog", "error", err)
	}
	return err
}

func (r IssueWorklogRepository) Destroy(worklogId uuid.UUID) error {
	err := r.DB.Delete(&IssueWorklogModel{}, worklogId).Error
	if err != nil {
		slog.Error("failed to destroy worklog", "error", err)
	}
	return err
}
