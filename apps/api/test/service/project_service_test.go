package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func buildProjectService(
	projectRepo domain.ProjectRepositoryInterface,
	issuePropertyRepo domain.ProjectIssuePropertyRepositoryInterface,
	trackerRepo domain.ProjectTrackerRepositoryInterface,
	trackerStatusRepo domain.ProjectTrackerStatusRepositoryInterface,
) domain.ProjectService {
	return domain.ProjectService{
		ProjectRepository:              projectRepo,
		ProjectIssuePropertyRepository: issuePropertyRepo,
		ProjectTrackerRepository:       trackerRepo,
		ProjectTrackerStatusRepository: trackerStatusRepo,
	}
}

func TestProjectService_All_ReturnsProjects(t *testing.T) {
	expected := &[]domain.ProjectModel{
		{ID: uuid.New(), Name: "Project A"},
		{ID: uuid.New(), Name: "Project B"},
	}
	projectRepo := &mockrepo.MockProjectRepository{
		AllFunc: func() (*[]domain.ProjectModel, error) {
			return expected, nil
		},
	}
	svc := buildProjectService(
		projectRepo,
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProjectService_All_ReturnsErrorOnFailure(t *testing.T) {
	repoErr := errors.New("db error")
	projectRepo := &mockrepo.MockProjectRepository{
		AllFunc: func() (*[]domain.ProjectModel, error) {
			return nil, repoErr
		},
	}
	svc := buildProjectService(
		projectRepo,
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	result, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
	assert.Nil(t, result)
}

func TestProjectService_Show_ReturnsProject(t *testing.T) {
	projectID := uuid.New()
	expected := &domain.ProjectModel{ID: projectID, Name: "Test Project"}
	projectRepo := &mockrepo.MockProjectRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			assert.Equal(t, projectID, id)
			return expected, nil
		},
	}
	svc := buildProjectService(
		projectRepo,
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	result, err := svc.Show(projectID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProjectService_Create_DelegatesCorrectly(t *testing.T) {
	var savedProject *domain.ProjectModel
	projectRepo := &mockrepo.MockProjectRepository{
		CreateFunc: func(project *domain.ProjectModel) error {
			savedProject = project
			return nil
		},
	}
	svc := buildProjectService(
		projectRepo,
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	project := &domain.ProjectModel{Name: "New Project", CodeName: "NP"}
	err := svc.Create(project)

	assert.NoError(t, err)
	assert.Equal(t, project, savedProject)
}

func TestProjectService_Update_DelegatesCorrectly(t *testing.T) {
	projectID := uuid.New()
	var updatedID uuid.UUID
	var updatedProject *domain.ProjectModel

	projectRepo := &mockrepo.MockProjectRepository{
		UpdateFunc: func(id uuid.UUID, project *domain.ProjectModel) error {
			updatedID = id
			updatedProject = project
			return nil
		},
	}
	svc := buildProjectService(
		projectRepo,
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	project := &domain.ProjectModel{Name: "Updated"}
	err := svc.Update(projectID, project)

	assert.NoError(t, err)
	assert.Equal(t, projectID, updatedID)
	assert.Equal(t, project, updatedProject)
}

func TestProjectService_Destroy_DelegatesCorrectly(t *testing.T) {
	projectID := uuid.New()
	var destroyedID uuid.UUID

	projectRepo := &mockrepo.MockProjectRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := buildProjectService(
		projectRepo,
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	err := svc.Destroy(projectID)

	assert.NoError(t, err)
	assert.Equal(t, projectID, destroyedID)
}

func TestProjectService_TrackersByProject_DelegatesCorrectly(t *testing.T) {
	projectID := uuid.New()
	expected := &[]domain.TrackerModel{
		{ID: uuid.New(), Name: "Bug"},
	}
	trackerRepo := &mockrepo.MockProjectTrackerRepository{
		FindTrackersByProjectIDFunc: func(id uuid.UUID) (*[]domain.TrackerModel, error) {
			assert.Equal(t, projectID, id)
			return expected, nil
		},
	}
	svc := buildProjectService(
		&mockrepo.MockProjectRepository{},
		&mockrepo.MockProjectIssuePropertyRepository{},
		trackerRepo,
		&mockrepo.MockProjectTrackerStatusRepository{},
	)

	result, err := svc.TrackersByProject(projectID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProjectService_StatusesByProject_DelegatesCorrectly(t *testing.T) {
	projectID := uuid.New()
	expected := &[]domain.StatusModel{
		{ID: uuid.New(), Name: "Open"},
	}
	statusRepo := &mockrepo.MockProjectTrackerStatusRepository{
		FindStatusesByProjectIDFunc: func(id uuid.UUID) (*[]domain.StatusModel, error) {
			assert.Equal(t, projectID, id)
			return expected, nil
		},
	}
	svc := buildProjectService(
		&mockrepo.MockProjectRepository{},
		&mockrepo.MockProjectIssuePropertyRepository{},
		&mockrepo.MockProjectTrackerRepository{},
		statusRepo,
	)

	result, err := svc.StatusesByProject(projectID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
