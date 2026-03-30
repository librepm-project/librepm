package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIssueService_All_ReturnsIssues(t *testing.T) {
	expected := &[]domain.IssueModel{
		{ID: uuid.New(), Summary: "Issue 1"},
		{ID: uuid.New(), Summary: "Issue 2"},
	}
	issueRepo := &mockrepo.MockIssueRepository{
		AllFunc: func() (*[]domain.IssueModel, error) {
			return expected, nil
		},
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: &mockrepo.MockProjectRepository{},
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestIssueService_All_ReturnsErrorOnFailure(t *testing.T) {
	repoErr := errors.New("db error")
	issueRepo := &mockrepo.MockIssueRepository{
		AllFunc: func() (*[]domain.IssueModel, error) {
			return nil, repoErr
		},
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: &mockrepo.MockProjectRepository{},
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	result, err := svc.All()

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestIssueService_Show_ReturnsIssue(t *testing.T) {
	issueID := uuid.New()
	expected := &domain.IssueModel{ID: issueID, Summary: "Test"}
	issueRepo := &mockrepo.MockIssueRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.IssueModel, error) {
			assert.Equal(t, issueID, id)
			return expected, nil
		},
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: &mockrepo.MockProjectRepository{},
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	result, err := svc.Show(issueID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestIssueService_Create_GeneratesKey(t *testing.T) {
	projectID := uuid.New()
	project := &domain.ProjectModel{
		ID:                 projectID,
		CodeName:           "TST",
		LastIssueKeyNumber: 1,
	}
	var createdIssue *domain.IssueModel

	projectRepo := &mockrepo.MockProjectRepository{
		IncrementLastIssueKeyNumberFunc: func(id uuid.UUID) error {
			assert.Equal(t, projectID, id)
			return nil
		},
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			return project, nil
		},
	}
	issueRepo := &mockrepo.MockIssueRepository{
		CreateFunc: func(issue *domain.IssueModel) error {
			createdIssue = issue
			return nil
		},
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: projectRepo,
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	issue := &domain.IssueModel{
		ProjectID: projectID,
		Summary:   "New issue",
	}
	err := svc.Create(issue)

	assert.NoError(t, err)
	assert.Equal(t, "TST-1", createdIssue.Key)
}

func TestIssueService_Create_AppliesDefaultTracker(t *testing.T) {
	projectID := uuid.New()
	trackerID := uuid.New()
	project := &domain.ProjectModel{
		ID:                 projectID,
		CodeName:           "TST",
		LastIssueKeyNumber: 5,
		DefaultTrackerID:   &trackerID,
	}

	projectRepo := &mockrepo.MockProjectRepository{
		IncrementLastIssueKeyNumberFunc: func(id uuid.UUID) error { return nil },
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			return project, nil
		},
	}
	issueRepo := &mockrepo.MockIssueRepository{
		CreateFunc: func(issue *domain.IssueModel) error { return nil },
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: projectRepo,
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	issue := &domain.IssueModel{
		ProjectID: projectID,
		Summary:   "New issue",
		TrackerID: uuid.Nil,
	}
	err := svc.Create(issue)

	assert.NoError(t, err)
	assert.Equal(t, trackerID, issue.TrackerID)
}

func TestIssueService_Create_AppliesDefaultStatus(t *testing.T) {
	projectID := uuid.New()
	statusID := uuid.New()
	project := &domain.ProjectModel{
		ID:                 projectID,
		CodeName:           "TST",
		LastIssueKeyNumber: 3,
		DefaultStatusID:    &statusID,
	}

	projectRepo := &mockrepo.MockProjectRepository{
		IncrementLastIssueKeyNumberFunc: func(id uuid.UUID) error { return nil },
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			return project, nil
		},
	}
	issueRepo := &mockrepo.MockIssueRepository{
		CreateFunc: func(issue *domain.IssueModel) error { return nil },
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: projectRepo,
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	issue := &domain.IssueModel{
		ProjectID: projectID,
		Summary:   "New issue",
		StatusID:  uuid.Nil,
	}
	err := svc.Create(issue)

	assert.NoError(t, err)
	assert.Equal(t, statusID, issue.StatusID)
}

func TestIssueService_Create_DoesNotOverrideExplicitTracker(t *testing.T) {
	projectID := uuid.New()
	defaultTrackerID := uuid.New()
	explicitTrackerID := uuid.New()
	project := &domain.ProjectModel{
		ID:                 projectID,
		CodeName:           "TST",
		LastIssueKeyNumber: 1,
		DefaultTrackerID:   &defaultTrackerID,
	}

	projectRepo := &mockrepo.MockProjectRepository{
		IncrementLastIssueKeyNumberFunc: func(id uuid.UUID) error { return nil },
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			return project, nil
		},
	}
	issueRepo := &mockrepo.MockIssueRepository{
		CreateFunc: func(issue *domain.IssueModel) error { return nil },
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: projectRepo,
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	issue := &domain.IssueModel{
		ProjectID: projectID,
		Summary:   "New issue",
		TrackerID: explicitTrackerID,
	}
	err := svc.Create(issue)

	assert.NoError(t, err)
	assert.Equal(t, explicitTrackerID, issue.TrackerID)
}

func TestIssueService_Create_FailsWhenIncrementFails(t *testing.T) {
	projectID := uuid.New()
	incrementErr := errors.New("increment failed")

	projectRepo := &mockrepo.MockProjectRepository{
		IncrementLastIssueKeyNumberFunc: func(id uuid.UUID) error { return incrementErr },
	}
	svc := domain.IssueService{
		IssueRepository:   &mockrepo.MockIssueRepository{},
		ProjectRepository: projectRepo,
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	issue := &domain.IssueModel{ProjectID: projectID, Summary: "New issue"}
	err := svc.Create(issue)

	assert.ErrorIs(t, err, incrementErr)
}

func TestIssueService_AllByFilterID_DelegatesCorrectly(t *testing.T) {
	filterID := uuid.New()
	projectID := uuid.New()
	conditions := []domain.FilterConditionModel{
		{Field: "project_id", Op: "eq", Value: projectID.String()},
	}
	filter := &domain.FilterModel{
		ID:               filterID,
		FilterConditions: conditions,
	}
	expectedIssues := &[]domain.IssueModel{
		{ID: uuid.New(), Summary: "Filtered issue"},
	}

	filterRepo := &mockrepo.MockFilterRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.FilterModel, error) {
			assert.Equal(t, filterID, id)
			return filter, nil
		},
	}
	issueRepo := &mockrepo.MockIssueRepository{
		AllByFilterFunc: func(conds []domain.FilterConditionModel) (*[]domain.IssueModel, error) {
			assert.Equal(t, conditions, conds)
			return expectedIssues, nil
		},
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: &mockrepo.MockProjectRepository{},
		FilterRepository:  filterRepo,
	}

	result, err := svc.AllByFilterID(filterID)

	assert.NoError(t, err)
	assert.Equal(t, expectedIssues, result)
}

func TestIssueService_Destroy_DelegatesCorrectly(t *testing.T) {
	issueID := uuid.New()
	var destroyedID uuid.UUID

	issueRepo := &mockrepo.MockIssueRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.IssueService{
		IssueRepository:   issueRepo,
		ProjectRepository: &mockrepo.MockProjectRepository{},
		FilterRepository:  &mockrepo.MockFilterRepository{},
	}

	err := svc.Destroy(issueID)

	assert.NoError(t, err)
	assert.Equal(t, issueID, destroyedID)
}
