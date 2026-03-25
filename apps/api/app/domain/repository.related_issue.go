package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RelatedIssueRepositoryInterface interface {
	Create(relation *RelatedIssueModel) error
	FindByIssue(issueID uuid.UUID, relationType *string) (*[]RelatedIssueModel, error)
	FindByID(relatedIssueID uuid.UUID) (*RelatedIssueModel, error)
	Delete(relatedIssueID uuid.UUID) error
}

type RelatedIssueRepository struct {
	DB *gorm.DB
}

func (r RelatedIssueRepository) Create(relation *RelatedIssueModel) error {
	return r.DB.Create(relation).Error
}

func (r RelatedIssueRepository) FindByIssue(issueID uuid.UUID, relationType *string) (*[]RelatedIssueModel, error) {
	var relations []RelatedIssueModel
	query := r.DB.Preload("IssueA").Preload("IssueB")

	if relationType != nil {
		query = query.Where("type = ?", *relationType)
	}

	query = query.Where("issue_a_id = ? OR issue_b_id = ?", issueID, issueID)

	if err := query.Find(&relations).Error; err != nil {
		return nil, err
	}
	return &relations, nil
}

func (r RelatedIssueRepository) FindByID(relatedIssueID uuid.UUID) (*RelatedIssueModel, error) {
	var relation RelatedIssueModel
	if err := r.DB.Preload("IssueA").Preload("IssueB").First(&relation, relatedIssueID).Error; err != nil {
		return nil, err
	}
	return &relation, nil
}

func (r RelatedIssueRepository) Delete(relatedIssueID uuid.UUID) error {
	return r.DB.Delete(&RelatedIssueModel{}, relatedIssueID).Error
}
