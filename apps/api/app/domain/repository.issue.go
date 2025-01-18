package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueRepositoryInterface interface {
	All() (*[]IssueModel, error)
	FindByID(issue_id uuid.UUID) (*IssueModel, error)
	Create(issue *IssueModel) error
	Update(issue_id uuid.UUID, issue *IssueModel) error
	Destroy(issue_id uuid.UUID) error
}

type IssueRepository struct {
	DB *gorm.DB
}

func (r IssueRepository) All() (*[]IssueModel, error) {
	var issues []IssueModel
	var err error
	query := r.DB.Select("issue.*")

	if err := query.Preload("Project").Preload("Tracker").Preload("Status").Find(&issues).Error; err != nil {
		fmt.Println(err)
	}
	return &issues, err
}

func (r IssueRepository) FindByID(issue_id uuid.UUID) (*IssueModel, error) {
	var issue IssueModel
	query := r.DB.Preload("Project").Preload("Tracker").Preload("Status").First(&issue, issue_id)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &issue, query.Error
}

func (r IssueRepository) Create(issue *IssueModel) error {
	query := r.DB.Create(&issue)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r IssueRepository) Update(issue_id uuid.UUID, issue *IssueModel) error {
	query := r.DB.Model(IssueModel{}).Where("id", issue_id).Updates(&issue)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r IssueRepository) Destroy(issue_id uuid.UUID) error {
	query := r.DB.Model(IssueModel{}).Delete(IssueModel{}, issue_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
