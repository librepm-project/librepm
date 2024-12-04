package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueRepositoryInterface interface {
	All() *[]IssueModel
	FindByID(issue_id uuid.UUID) *IssueModel
	Create(issue *IssueModel)
	Update(issue_id uuid.UUID, issue *IssueModel)
	Destroy(issue_id uuid.UUID)
}

type IssueRepository struct {
	DB *gorm.DB
}

func (r IssueRepository) All() *[]IssueModel {
	var issues []IssueModel
	query := r.DB.Select("issue.*")

	if err := query.Preload("Project").Preload("Tracker").Preload("Status").Find(&issues).Error; err != nil {
		fmt.Println(err)
	}
	return &issues
}

func (r IssueRepository) FindByID(issue_id uuid.UUID) *IssueModel {
	var issue IssueModel
	fmt.Println(issue_id)
	err := r.DB.Preload("Project").Preload("Tracker").Preload("Status").First(&issue, issue_id)

	if err != nil {
		fmt.Println(err)
	}
	return &issue
}

func (r IssueRepository) Create(issue *IssueModel) {
	r.DB.Create(&issue)
}

func (r IssueRepository) Update(issue_id uuid.UUID, issue *IssueModel) {
	r.DB.Model(IssueModel{}).Where("id", issue_id).Updates(&issue)
}

func (r IssueRepository) Destroy(issue_id uuid.UUID) {
	r.DB.Model(IssueModel{}).Delete(IssueModel{}, issue_id)
}
