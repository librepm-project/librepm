package service_test

import (
	"errors"
	"testing"
	"time"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWorklogService_AllByIssue_ReturnsWorklogs(t *testing.T) {
	issueID := uuid.New()
	expected := &[]domain.IssueWorklogModel{
		{ID: uuid.New(), Seconds: 3600, Comment: "Fixed it"},
	}
	repo := &mockrepo.MockIssueWorklogRepository{
		AllByIssueFunc: func(id uuid.UUID) (*[]domain.IssueWorklogModel, error) {
			assert.Equal(t, issueID, id)
			return expected, nil
		},
	}
	svc := domain.IssueWorklogService{IssueWorklogRepository: repo}

	result, err := svc.AllByIssue(issueID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestWorklogService_AllByIssue_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockIssueWorklogRepository{
		AllByIssueFunc: func(_ uuid.UUID) (*[]domain.IssueWorklogModel, error) {
			return nil, repoErr
		},
	}
	svc := domain.IssueWorklogService{IssueWorklogRepository: repo}

	_, err := svc.AllByIssue(uuid.New())

	assert.ErrorIs(t, err, repoErr)
}

func TestWorklogService_Show_ReturnsWorklog(t *testing.T) {
	worklogID := uuid.New()
	expected := &domain.IssueWorklogModel{ID: worklogID, Seconds: 1800}
	repo := &mockrepo.MockIssueWorklogRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.IssueWorklogModel, error) {
			assert.Equal(t, worklogID, id)
			return expected, nil
		},
	}
	svc := domain.IssueWorklogService{IssueWorklogRepository: repo}

	result, err := svc.Show(worklogID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestWorklogService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.IssueWorklogModel
	repo := &mockrepo.MockIssueWorklogRepository{
		CreateFunc: func(worklog *domain.IssueWorklogModel) error {
			saved = worklog
			return nil
		},
	}
	svc := domain.IssueWorklogService{IssueWorklogRepository: repo}

	worklog := &domain.IssueWorklogModel{
		IssueID:  uuid.New(),
		Seconds:  7200,
		Comment:  "Worked on feature",
		LoggedAt: time.Now(),
	}
	err := svc.Create(worklog)

	assert.NoError(t, err)
	assert.Equal(t, worklog, saved)
}

func TestWorklogService_Update_DelegatesCorrectly(t *testing.T) {
	worklogID := uuid.New()
	var updatedID uuid.UUID
	repo := &mockrepo.MockIssueWorklogRepository{
		UpdateFunc: func(id uuid.UUID, worklog *domain.IssueWorklogModel) error {
			updatedID = id
			return nil
		},
	}
	svc := domain.IssueWorklogService{IssueWorklogRepository: repo}

	err := svc.Update(worklogID, &domain.IssueWorklogModel{Seconds: 3600})

	assert.NoError(t, err)
	assert.Equal(t, worklogID, updatedID)
}

func TestWorklogService_Destroy_DelegatesCorrectly(t *testing.T) {
	worklogID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockIssueWorklogRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.IssueWorklogService{IssueWorklogRepository: repo}

	err := svc.Destroy(worklogID)

	assert.NoError(t, err)
	assert.Equal(t, worklogID, destroyedID)
}
