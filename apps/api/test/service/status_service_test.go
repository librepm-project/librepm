package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStatusService_All_ReturnsStatuses(t *testing.T) {
	expected := &[]domain.StatusModel{
		{ID: uuid.New(), Name: "Open", Color: "#00ff00"},
	}
	repo := &mockrepo.MockStatusRepository{
		AllFunc: func() (*[]domain.StatusModel, error) { return expected, nil },
	}
	svc := domain.StatusService{StatusRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestStatusService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockStatusRepository{
		AllFunc: func() (*[]domain.StatusModel, error) { return nil, repoErr },
	}
	svc := domain.StatusService{StatusRepository: repo}

	result, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
	assert.Nil(t, result)
}

func TestStatusService_Show_ReturnsStatus(t *testing.T) {
	statusID := uuid.New()
	expected := &domain.StatusModel{ID: statusID, Name: "Closed"}
	repo := &mockrepo.MockStatusRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.StatusModel, error) {
			assert.Equal(t, statusID, id)
			return expected, nil
		},
	}
	svc := domain.StatusService{StatusRepository: repo}

	result, err := svc.Show(statusID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestStatusService_Create_DelegatesCorrectly(t *testing.T) {
	var savedStatus *domain.StatusModel
	repo := &mockrepo.MockStatusRepository{
		CreateFunc: func(status *domain.StatusModel) error {
			savedStatus = status
			return nil
		},
	}
	svc := domain.StatusService{StatusRepository: repo}

	status := &domain.StatusModel{Name: "In Progress", Color: "#ffff00"}
	err := svc.Create(status)

	assert.NoError(t, err)
	assert.Equal(t, status, savedStatus)
}

func TestStatusService_Update_DelegatesCorrectly(t *testing.T) {
	statusID := uuid.New()
	var updatedID uuid.UUID
	repo := &mockrepo.MockStatusRepository{
		UpdateFunc: func(id uuid.UUID, status *domain.StatusModel) error {
			updatedID = id
			return nil
		},
	}
	svc := domain.StatusService{StatusRepository: repo}

	err := svc.Update(statusID, &domain.StatusModel{Name: "Updated"})

	assert.NoError(t, err)
	assert.Equal(t, statusID, updatedID)
}

func TestStatusService_Destroy_DelegatesCorrectly(t *testing.T) {
	statusID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockStatusRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.StatusService{StatusRepository: repo}

	err := svc.Destroy(statusID)

	assert.NoError(t, err)
	assert.Equal(t, statusID, destroyedID)
}
