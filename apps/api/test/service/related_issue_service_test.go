package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRelatedIssueService_Create_SelfReferenceReturnsError(t *testing.T) {
	issueID := uuid.New()
	svc := domain.RelatedIssueService{
		RelatedIssueRepository: &mockrepo.MockRelatedIssueRepository{},
		IssueRepository:        &mockrepo.MockIssueRepository{},
	}

	result, err := svc.Create(issueID, issueID, "blocks")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "self-referencing")
	assert.Nil(t, result)
}

func TestRelatedIssueService_Create_IssueANotFoundReturnsError(t *testing.T) {
	issueAID := uuid.New()
	issueBID := uuid.New()
	repoErr := errors.New("not found")

	svc := domain.RelatedIssueService{
		RelatedIssueRepository: &mockrepo.MockRelatedIssueRepository{},
		IssueRepository: &mockrepo.MockIssueRepository{
			FindByIDFunc: func(id uuid.UUID) (*domain.IssueModel, error) {
				if id == issueAID {
					return nil, repoErr
				}
				return &domain.IssueModel{ID: id}, nil
			},
		},
	}

	result, err := svc.Create(issueAID, issueBID, "blocks")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "issue A not found")
	assert.Nil(t, result)
}

func TestRelatedIssueService_Create_IssueBNotFoundReturnsError(t *testing.T) {
	issueAID := uuid.New()
	issueBID := uuid.New()
	repoErr := errors.New("not found")

	svc := domain.RelatedIssueService{
		RelatedIssueRepository: &mockrepo.MockRelatedIssueRepository{},
		IssueRepository: &mockrepo.MockIssueRepository{
			FindByIDFunc: func(id uuid.UUID) (*domain.IssueModel, error) {
				if id == issueBID {
					return nil, repoErr
				}
				return &domain.IssueModel{ID: id}, nil
			},
		},
	}

	result, err := svc.Create(issueAID, issueBID, "blocks")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "issue B not found")
	assert.Nil(t, result)
}

func TestRelatedIssueService_Create_CreatesAndReturnsRelation(t *testing.T) {
	issueAID := uuid.New()
	issueBID := uuid.New()
	relID := uuid.New()

	svc := domain.RelatedIssueService{
		RelatedIssueRepository: &mockrepo.MockRelatedIssueRepository{
			CreateFunc: func(rel *domain.RelatedIssueModel) error { return nil },
			FindByIDFunc: func(id uuid.UUID) (*domain.RelatedIssueModel, error) {
				return &domain.RelatedIssueModel{ID: relID, IssueAID: issueAID, IssueBID: issueBID, Type: "blocks"}, nil
			},
		},
		IssueRepository: &mockrepo.MockIssueRepository{
			FindByIDFunc: func(id uuid.UUID) (*domain.IssueModel, error) {
				return &domain.IssueModel{ID: id}, nil
			},
		},
	}

	result, err := svc.Create(issueAID, issueBID, "blocks")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, issueAID, result.IssueAID)
	assert.Equal(t, issueBID, result.IssueBID)
	assert.Equal(t, "blocks", result.Type)
}

func TestRelatedIssueService_GetRelatedIssues_DelegatesCorrectly(t *testing.T) {
	issueID := uuid.New()
	relationType := "blocks"
	expected := &[]domain.RelatedIssueModel{
		{ID: uuid.New(), IssueAID: issueID, Type: relationType},
	}
	svc := domain.RelatedIssueService{
		RelatedIssueRepository: &mockrepo.MockRelatedIssueRepository{
			FindByIssueFunc: func(id uuid.UUID, rt *string) (*[]domain.RelatedIssueModel, error) {
				assert.Equal(t, issueID, id)
				assert.Equal(t, relationType, *rt)
				return expected, nil
			},
		},
		IssueRepository: &mockrepo.MockIssueRepository{},
	}

	result, err := svc.GetRelatedIssues(issueID, &relationType)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestRelatedIssueService_Delete_DelegatesCorrectly(t *testing.T) {
	relID := uuid.New()
	var deletedID uuid.UUID
	svc := domain.RelatedIssueService{
		RelatedIssueRepository: &mockrepo.MockRelatedIssueRepository{
			DeleteFunc: func(id uuid.UUID) error {
				deletedID = id
				return nil
			},
		},
		IssueRepository: &mockrepo.MockIssueRepository{},
	}

	err := svc.Delete(relID)

	assert.NoError(t, err)
	assert.Equal(t, relID, deletedID)
}
